package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func loadDotEnv() error {
    envFiles := []string{
        ".env.local",
        ".env.dev.local",
        ".env.stating.local",
        ".env.prod.local",
        ".env.test.local",
    }

    var loadedFile string

    for _, filename := range envFiles {
        
        if _, err := os.Stat(filename); err == nil {
            if err := godotenv.Load(filename); err != nil {
                return fmt.Errorf("error loading env file %s: %v", filename, err)
            }
            
            loadedFile = filename
            break
        }
    }


    if loadedFile != "" {
        log.Printf("Loaded environment configuration from: %s", loadedFile)
    } else {
        log.Println("No environment variables file found. Using environment variables or defaults.")
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
