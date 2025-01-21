package dbutils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	_ "url_shortener/models"
)

func ConnectDB() (*sql.DB, error) {

	// Retrieve environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	// Connect to MySQL database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

func InitDB(db *sql.DB) error {
	// Initialize the table if not exists
	query := `
		CREATE TABLE IF NOT EXISTS url_mappings (
			id BIGINT AUTO_INCREMENT PRIMARY KEY,
			original_url TEXT NOT NULL,
			short_url VARCHAR(255) NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err := db.Exec(query)
	return err
}

func AddURL(db *sql.DB, original, short string) (int64, error) {
	query := `INSERT INTO url_mappings(original_url, short_url, created_at) VALUES (?, ?, NOW())`
	result, err := db.Exec(query, original, short)
	if err != nil {
		return 0, fmt.Errorf("addURL: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addURL LastInsertId: %v", err)
	}
	return id, nil
}

func CheckIfURLExists(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM url_mappings WHERE short_url=?`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("shortURL %s does not exist", shortURL)
	}
	return originalURL, err
}
