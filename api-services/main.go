package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

type URLMappings struct {
	ID          int64     `json:"id"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

func connectDB() (*sql.DB, error) {
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

func addURL(db *sql.DB, original URLMappings) (int64, error) {

	query := `INSERT INTO url_mappings(original_url, short_url, created_at) VALUES (?, ?, ?)`
	result, err := db.Exec(query, original.OriginalURL, original.ShortURL, original.CreatedAt)

	if err != nil {
		return 0, fmt.Errorf("addURL: %v", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addURL LastInsertId: %v", err)
	}
	return id, nil
}

func handlerPing(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}

func urlPOST(c *gin.Context) {
	var request URLMappings

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//checking if the provided URL is not empty

	if request.OriginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Original URL cannot be empty"})
		return
	}

	//connecting to the database
	db, err := connectDB()

	//check to ensure that database connectivity is working fine.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database."})
	}
	defer db.Close()

	var shortURL string

	for {
		shortURL = urlGenerator()
		var exists int

		err = db.QueryRow("SELECT COUNT(*) FROM url_mappings where short_url = ?", shortURL).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate uniqueness of short URL"})
		}
		if exists == 0 {
			break
		}
	}

	newMapping := URLMappings{
		OriginalURL: request.OriginalURL,
		ShortURL:    shortURL,
		CreatedAt:   time.Now(),
	}

	_, err = addURL(db, newMapping)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Short URL created successfully",
		"shortURL": shortURL,
	})

	//fmt.Println("Received URL:", originalUrl.OriginalURL)
	return
}

func urlGenerator() string {
	random := rand.IntN(9000) + 1000

	prefix := "simple-url"
	fmt.Println(random)

	shortURL := fmt.Sprintf("%s-%d", prefix, random)
	return shortURL
}

func initDB(db *sql.DB) error {
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

func main() {

	db, err := connectDB()

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/ping", handlerPing)

	r.POST("/shorten", urlPOST)

	r.Run("localhost:8080")

}
