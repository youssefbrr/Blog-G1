package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
)

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

// loadDotEnv attempts to load .env file from current, parent, or grandparent directories
func loadDotEnv() error {
    searchDirs := []string{
        ".",
        "..",
        "../..",
    }

    envFiles := []string{
        ".env",
        "app.env",
        "config.env",
    }

    var loadedFile string

    for _, dir := range searchDirs {
        for _, filename := range envFiles {
            filepath := filepath.Join(dir, filename)
            
            if _, err := os.Stat(filepath); err == nil {
                if err := godotenv.Load(filepath); err != nil {
                    return fmt.Errorf("error loading env file %s: %v", filepath, err)
                }
                
                loadedFile = filepath
                break
            }
        }

        if loadedFile != "" {
            break
        }
    }

    if loadedFile != "" {
        log.Printf("Loaded environment configuration from: %s", loadedFile)
    } else {
        log.Println("No .env file found. Using environment variables or defaults.")
    }

    return nil
}

func getEnvWithDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getDSNFromEnv() string {
    dbUser := getEnvWithDefault("MYSQL_USER", "root")
    dbPass := os.Getenv("MYSQL_PASSWORD")
    dbHost := getEnvWithDefault("MYSQL_HOST", "127.0.0.1")
    dbPort := getEnvWithDefault("MYSQL_PORT", "8080")
    dbName := getEnvWithDefault("MYSQL_DATABASE", "blogdb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        dbUser, dbPass, dbHost, dbPort, dbName)
    fmt.Println(dsn);
    return dsn
}
