package database

import (
	"fmt"
	"learn_service/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
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