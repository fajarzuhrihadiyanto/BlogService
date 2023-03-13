package utils

import (
	"MyBlog/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectToDB
// This function is used to connect to database with connection string
func ConnectToDB() (*gorm.DB, error) {
	// Get connection string from environment variable
	connectionString := GetEnvVariable("DATABASE_URL")

	// Open database connection
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	err = db.AutoMigrate(&models.User{}, &models.Article{})

	// Return database instance and the error if any
	return db, err
}
