package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

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
