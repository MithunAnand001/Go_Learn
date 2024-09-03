package database

import (
	"fmt"
	"learn_service/constant"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    // Connect to the database
    db, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=%s",
	constant.DBUser, constant.DBName, constant.DBPassword, constant.DBHost, constant.DBPort, constant.DBSSLMode)), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // Ping the database to verify the connection
    err = sqlDB.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

// CloseDB closes the database connection.
func CloseDB(db *gorm.DB) {
	// Get the underlying *sql.DB object from the *gorm.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Error getting DB from gorm.DB:", err)
		return
	}

	// Close the database connection
	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing the database connection:", err)
	}
}