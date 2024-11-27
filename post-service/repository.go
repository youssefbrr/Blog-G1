package main

import (
    "database/sql"
    "fmt"
    "time"

    "github.com/go-sql-driver/mysql"
)

type PostRepository struct {
    db *sql.DB
}

func NewPostRepository(dataSourceName string) (*PostRepository, error) {
    _, err := mysql.ParseDSN(dataSourceName)
    if err != nil {
        return nil, fmt.Errorf("invalid MySQL DSN: %v", err)
    }

    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Post (
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

func (r *PostRepository) Create(post Post) error {
    query := `
        INSERT INTO Post
        (id, title, content, author, created_at) 
        VALUES (?, ?, ?, ?, ?)
    `
    _, err := r.db.Exec(query, post.ID, post.Title, post.Content, 
        post.Author, post.CreatedAt)
    return err
}

func (r *PostRepository) Get(id string) (Post, error) {
    query := `
        SELECT id, title, content, author, created_at, updated_at, deleted_at 
        FROM Post
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

    if updatedAt.Valid {
        post.UpdatedAt = &updatedAt.Time
    }
    if deletedAt.Valid {
        post.DeletedAt = &deletedAt.Time
    }

    return post, nil
}

func (r *PostRepository) Update(id string, post Post) error {
    query := `
        UPDATE Post
        SET title = ?, content = ?, updated_at = ? 
        WHERE id = ? AND deleted_at IS NULL
    `
    result, err := r.db.Exec(query, post.Title, post.Content, 
        post.UpdatedAt, id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    if rowsAffected == 0 {
        return fmt.Errorf("post not found or already deleted")
    }

    return nil
}

// soft deletes a post
func (r *PostRepository) Delete(id string) error {
    query := `
        UPDATE Post
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
        FROM Post
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

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return posts, nil
}

// Close closes the database connection
func (r *PostRepository) Close() error {
    return r.db.Close()
}
