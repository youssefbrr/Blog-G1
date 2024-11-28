package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    loadDotEnv()
    
    dsn := getDSNFromEnv()

    repo, err := NewPostRepository(dsn)
    if err != nil {
        log.Fatalf("Failed to create repository: %v", err)
    }
    defer repo.Close()

    // Set Gin to production mode
    gin.SetMode(gin.ReleaseMode)

    // Create router
    r := gin.Default()

r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    postHandler := NewPostHandler(repo)

    v1 := r.Group("/api/v1/posts")
    {
        v1.GET("/healthcheck", Check)
        v1.POST("", postHandler.CreatePost)
        v1.GET("", postHandler.ListPosts)
        v1.GET("/:id", postHandler.GetPost)
        v1.PUT("/:id", postHandler.UpdatePost)
        v1.DELETE("/:id", postHandler.DeletePost)
    }

    port := getEnvWithDefault("PMA_PORT", "8080")
    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}