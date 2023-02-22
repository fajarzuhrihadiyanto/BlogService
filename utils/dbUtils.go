package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// ConnectToDB
// This function is used to connect to database with connection string
func ConnectToDB() (*sql.DB, error) {
	// Get connection string from environment variable
	connectionString := GetEnvVariable("DATABASE_URL")

	// Initiate connection to mysql server using connection string
	db, err := sql.Open(`mysql`, connectionString)

	// Return the error if any
	return db, err
}
