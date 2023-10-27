package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/layardaputra/govtech-catalog-test-project/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()
	// Database connection settings
	dbURL := cfg.DatabaseURL

	// Open a connection to the database
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Execute SQL scripts
	err = executeSQLMigration(db)
	if err != nil {
		log.Fatalf("Error executing DDL: %v", err)
	}
	log.Println("DDL executed successfully")
}

func executeSQLMigration(db *sqlx.DB) error {
	sql := `
	CREATE TABLE product (
		id BIGSERIAL PRIMARY KEY,
		sku VARCHAR(255),
		title VARCHAR(255),
		description TEXT,
		category VARCHAR(255),
		etalase VARCHAR(255),
		images JSONB,
		weight NUMERIC,
		price NUMERIC,
		created_at TIMESTAMPTZ,
		updated_at TIMESTAMPTZ
	);
	
	CREATE INDEX idx_product_sku on product (sku);
	CREATE INDEX idx_product_title on product (title);
	CREATE INDEX idx_product_category on product (category);
	CREATE INDEX idx_product_etalase on product (etalase);
	CREATE INDEX idx_product_created_at on product (created_at);
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback() // Rollback if an error occurs

	_, err = tx.Exec(sql)
	if err != nil {
		return err
	}

	return tx.Commit()
}
