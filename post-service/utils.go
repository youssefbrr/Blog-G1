package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/joho/godotenv"
)

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

// loadDotEnv attempts to load .env file from current, parent, or grandparent directories
func loadDotEnv() error {
    // Directories to search for .env file (in order)
    searchDirs := []string{
        ".",          // Current directory
        "..",         // Parent directory
        "../..",      // Grandparent directory
    }

    // Potential .env file names
    envFiles := []string{
        ".env",
        "app.env",
        "config.env",
    }

    // Track loaded file for logging
    var loadedFile string

    // Attempt to load .env file
    for _, dir := range searchDirs {
        for _, filename := range envFiles {
            // Construct full path
            filepath := filepath.Join(dir, filename)
            
            // Check if file exists
            if _, err := os.Stat(filepath); err == nil {
                // Try to load the file
                if err := godotenv.Load(filepath); err != nil {
                    return fmt.Errorf("error loading env file %s: %v", filepath, err)
                }
                
                loadedFile = filepath
                break
            }
        }

        // If a file was loaded, stop searching
        if loadedFile != "" {
            break
        }
    }

    // Log which file was loaded (or not)
    if loadedFile != "" {
        log.Printf("Loaded environment configuration from: %s", loadedFile)
    } else {
        log.Println("No .env file found. Using environment variables or defaults.")
    }

    return nil
}

// getEnvWithDefault retrieves an environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getDSNFromEnv() string {
    // Read database configuration from environment variables
    dbUser := getEnvWithDefault("MYSQL_USER", "root")
    dbPass := os.Getenv("MYSQL_PASSWORD")
    dbHost := getEnvWithDefault("MYSQL_HOST", "127.0.0.1")
    dbPort := getEnvWithDefault("MYSQL_PORT", "8080")
    dbName := getEnvWithDefault("MYSQL_DATABASE", "blogdb")

    // Construct DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
        dbUser, dbPass, dbHost, dbPort, dbName)
    fmt.Println(dsn);
    return dsn
}
