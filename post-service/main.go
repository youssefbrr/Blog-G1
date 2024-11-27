package main

import (
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    // Load configuration from .env
    loadDotEnv()
    
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
    port := getEnvWithDefault("PMA_PORT", "8080")
    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}