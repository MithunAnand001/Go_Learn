package constant

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

var (
    DBUser string
    DBName       string
    DBPassword   string
    DBHost       string
    DBPort       string
    DBSSLMode    string
)

func init() {
    // Load environment variables from the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Assign values from environment variables
    DBUser = os.Getenv("DB_USER")
    DBName = os.Getenv("DB_NAME")
    DBPassword = os.Getenv("DB_PASSWORD")
    DBHost = os.Getenv("DB_HOST")
    DBPort = os.Getenv("DB_PORT")
}