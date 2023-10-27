package database

import (
	"log"

	// Import the SQLx and the PostgreSQL driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitDB initializes the database connection.
func InitDB(dataSourceName string) (*sqlx.DB, error) {
	// Open a connection to the database
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping the database to ensure the connection is established
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	return db, nil
}
