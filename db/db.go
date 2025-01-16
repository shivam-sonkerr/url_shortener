package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/url_mappings")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(db *sql.DB) error {
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
