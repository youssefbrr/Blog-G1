package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// Post represents the post data structure
type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Author    string     `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// PostRepository handles database operations for posts
type PostRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new PostRepository
func NewPostRepository(dataSourceName string) (*PostRepository, error) {
	// Configure MySQL connection
	_, err := mysql.ParseDSN(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("invalid MySQL DSN: %v", err)
	}

	// Open database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Create posts table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id VARCHAR(36) PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			author VARCHAR(50) NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME,
			deleted_at DATETIME
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create posts table: %v", err)
	}

	return &PostRepository{db: db}, nil
}

// Create adds a new post to the database
func (r *PostRepository) Create(post Post) error {
	query := `
		INSERT INTO posts 
		(id, title, content, author, created_at) 
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, post.ID, post.Title, post.Content, 
		post.Author, post.CreatedAt)
	return err
}

// Get retrieves a post by ID
func (r *PostRepository) Get(id string) (Post, error) {
	query := `
		SELECT id, title, content, author, created_at, updated_at, deleted_at 
		FROM posts 
		WHERE id = ? AND deleted_at IS NULL
	`
	var post Post
	var updatedAt, deletedAt sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&post.ID, &post.Title, &post.Content, &post.Author, 
		&post.CreatedAt, &updatedAt, &deletedAt,
	)
	if err == sql.ErrNoRows {
		return Post{}, fmt.Errorf("post not found")
	}
	if err != nil {
		return Post{}, err
	}

	// Handle nullable timestamp fields
	if updatedAt.Valid {
		post.UpdatedAt = &updatedAt.Time
	}
	if deletedAt.Valid {
		post.DeletedAt = &deletedAt.Time
	}

	return post, nil
}

// Update modifies an existing post
func (r *PostRepository) Update(id string, post Post) error {
	query := `
		UPDATE posts 
		SET title = ?, content = ?, updated_at = ? 
		WHERE id = ? AND deleted_at IS NULL
	`
	result, err := r.db.Exec(query, post.Title, post.Content, 
		post.UpdatedAt, id)
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post not found or already deleted")
	}

	return nil
}

// Delete soft deletes a post
func (r *PostRepository) Delete(id string) error {
	query := `
		UPDATE posts 
		SET deleted_at = ? 
		WHERE id = ? AND deleted_at IS NULL
	`
	now := time.Now()
	result, err := r.db.Exec(query, now, id)
	if err != nil {
		return err
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post not found or already deleted")
	}

	return nil
}

// List returns all non-deleted posts
func (r *PostRepository) List() ([]Post, error) {
	query := `
		SELECT id, title, content, author, created_at, updated_at 
		FROM posts 
		WHERE deleted_at IS NULL
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var updatedAt sql.NullTime

		if err := rows.Scan(
			&post.ID, &post.Title, &post.Content, &post.Author, 
			&post.CreatedAt, &updatedAt,
		); err != nil {
			return nil, err
		}

		if updatedAt.Valid {
			post.UpdatedAt = &updatedAt.Time
		}

		posts = append(posts, post)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Close closes the database connection
func (r *PostRepository) Close() error {
	return r.db.Close()
}

// Validation function
func validatePost(post Post) error {
	if post.Title == "" {
		return fmt.Errorf("title is required")
	}
	if post.Content == "" {
		return fmt.Errorf("content is required")
	}
	if post.Author == "" {
		return fmt.Errorf("author is required")
	}
	if len(post.Title) > 255 {
		return fmt.Errorf("title cannot exceed 255 characters")
	}
	if len(post.Author) > 50 {
		return fmt.Errorf("author cannot exceed 50 characters")
	}
	return nil
}

// PostHandler manages post-related HTTP operations
type PostHandler struct {
	repo *PostRepository
}

// NewPostHandler creates a new PostHandler
func NewPostHandler(repo *PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

// CreatePost handles POST requests to create a new post
func (ph *PostHandler) CreatePost(c *gin.Context) {
	var newPost Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the post
	if err := validatePost(newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a unique ID and timestamps
	newPost.ID = uuid.New().String()
	newPost.CreatedAt = time.Now()

	// Store the new post
	if err := ph.repo.Create(newPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)
}

// GetPost handles GET requests to retrieve a post by ID
func (ph *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	
	post, err := ph.repo.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, post)
}

// UpdatePost handles PUT requests to update an existing post
func (ph *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var updatedData Post
	
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the updated post
	if err := validatePost(updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the existing post to ensure it exists
	existingPost, err := ph.repo.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Preserve original metadata
	updatedData.ID = existingPost.ID
	updatedData.CreatedAt = existingPost.CreatedAt
	
	// Update timestamp
	now := time.Now()
	updatedData.UpdatedAt = &now

	// Save the updated post
	if err := ph.repo.Update(id, updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedData)
}

// DeletePost handles DELETE requests to remove a post
func (ph *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	
	if err := ph.repo.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// ListPosts handles GET requests to list all posts
func (ph *PostHandler) ListPosts(c *gin.Context) {
	postList, err := ph.repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, postList)
}

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, "{'status': 200, 'message': 'service is running!'}")
}

func loadConfiguration() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using environment variables or defaults.")
	}
}

func getDSNFromEnv() string {
	// Read database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Default values if not set
	if dbUser == "" {
		dbUser = "root"
	}
	if dbPort == "" {
		dbPort = "3306"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbName == "" {
		dbName = "blogdb"
	}

	// Construct DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		dbUser, dbPass, dbHost, dbPort, dbName)
	
	return dsn
}

func main() {
	// MySQL connection string format
	// Format: username:password@tcp(hostname:port)/database_name

	// Load configuration from .env
	loadConfiguration()
	
	// Get DSN from environment variables
	dsn := getDSNFromEnv()

	// Create post repository
	repo, err := NewPostRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}
	defer repo.Close()

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.Default()

	// Create post handler
	postHandler := NewPostHandler(repo)

	// Define routes
	v1 := r.Group("/api/v1/posts")
	{
		v1.GET("/healthcheck", Check)
		v1.POST("", postHandler.CreatePost)
		v1.GET("", postHandler.ListPosts)
		v1.GET("/:id", postHandler.GetPost)
		v1.PUT("/:id", postHandler.UpdatePost)
		v1.DELETE("/:id", postHandler.DeletePost)
	}

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}