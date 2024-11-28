package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostHandler struct {
    repo *PostRepository
}

func NewPostHandler(repo *PostRepository) *PostHandler {
    return &PostHandler{repo: repo}
}

func (ph *PostHandler) CreatePost(c *gin.Context) {
    var newPost Post
    if err := c.ShouldBindJSON(&newPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validatePostCreateBody(newPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newPost.ID = uuid.New().String()
    newPost.CreatedAt = time.Now()

    if err := ph.repo.Create(newPost); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newPost)
}

func (ph *PostHandler) GetPost(c *gin.Context) {
    id := c.Param("id")
    
    post, err := ph.repo.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusOK, post)
}

func (ph *PostHandler) UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var updatedData Post
    
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validatePostUpdateBody(updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    existingPost, err := ph.repo.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    updatedData.ID = existingPost.ID
    updatedData.CreatedAt = existingPost.CreatedAt
    
    now := time.Now()
    updatedData.UpdatedAt = &now

    if err := ph.repo.Update(id, updatedData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedData)
}

func (ph *PostHandler) PartialUpdatePost(c *gin.Context) {
    id := c.Param("id")

    var updatedData Post
    
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validatePostPartialUpdateBody(updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    existingPost, err := ph.repo.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    updatedData.ID = existingPost.ID
    if updatedData.Title == "" {
        updatedData.Title = existingPost.Title
    }
    if updatedData.Content == "" {
        updatedData.Content = existingPost.Content
    }
    updatedData.Author = existingPost.Author
    updatedData.CreatedAt = existingPost.CreatedAt

    now := time.Now()
    updatedData.UpdatedAt = &now

    if err := ph.repo.Update(id, updatedData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedData)
}

func (ph *PostHandler) DeletePost(c *gin.Context) {
    id := c.Param("id")
    
    if err := ph.repo.Delete(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    
    c.Status(http.StatusNoContent)
}

func (ph *PostHandler) ListPosts(c *gin.Context) {
    postList, err := ph.repo.List()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, postList)
}

func Check(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "message": "service is running!",
    })
}
