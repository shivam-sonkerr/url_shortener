package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"url_shortener/dbutils"
	"url_shortener/models"
)

func handlerPing(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}

func redirectURLHandler(c *gin.Context) {

	db, err := dbutils.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close() // Close the database connection after the handler completes

	shortURL := c.Param("shortURL")

	fmt.Println("Redirecting for shortURL:", shortURL) // log the shortURL to confirm

	originalURL, err := dbutils.CheckIfURLExists(db, shortURL)

	//if err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "NOT FOUND"})
	//	return
	//}

	if originalURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "shortURL has no valid mapping"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func urlPOST(c *gin.Context) {
	var request models.URLMappings

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Validate the Original URL
	if !isValidURL(request.OriginalURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	//checking if the provided URL is not empty

	if request.OriginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Original URL cannot be empty"})
		return
	}

	//connecting to the database
	db, err := dbutils.ConnectDB()

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

	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Short URL created successfully",
		"shortURL": shortURL,
	})

	fmt.Printf("New URL mapping added: %s -> %s\n", request.OriginalURL, shortURL)

	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://localhost:8080/redirect/%s", shortURL))
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

//func initDB(db *sql.DB) error {
//	query := `
//       CREATE TABLE IF NOT EXISTS url_mappings (
//           id BIGINT AUTO_INCREMENT PRIMARY KEY,
//           original_url TEXT NOT NULL,
//           short_url VARCHAR(255) NOT NULL UNIQUE,
//           created_at DATETIME DEFAULT CURRENT_TIMESTAMP
//       );
//   `
//	_, err := db.Exec(query)
//	return err
//}

func shortenAndRedirect(c *gin.Context) {
	var request models.URLMappings

	// Parse the request body
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate the provided URL
	if request.OriginalURL == "" || !isValidURL(request.OriginalURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or empty URL"})
		return
	}

	// Connect to the database
	db, err := dbutils.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	// Generate a unique short URL
	var shortURL string
	for {
		shortURL = urlGenerator()
		var exists int
		err = db.QueryRow("SELECT COUNT(*) FROM url_mappings WHERE short_url = ?", shortURL).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		if exists == 0 {
			break
		}
	}

	// Save the URL mapping in the database
	//newMapping := models.URLMappings{
	//	OriginalURL: request.OriginalURL,
	//	ShortURL:    shortURL,
	//	CreatedAt:   time.Now(),
	//}
	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	// Redirect the user to the original URL
	c.Redirect(http.StatusMovedPermanently, request.OriginalURL)
}

func main() {

	db, err := dbutils.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	err = dbutils.InitDB(db)
	if err != nil {
		log.Fatal("Failed to initialize database", err)
	}

	r := gin.Default()

	r.GET("/ping", handlerPing)

	r.POST("/shorten", urlPOST)

	r.POST("/shorten-and-redirect", shortenAndRedirect)

	r.GET("/redirect/:shortURL", redirectURLHandler)

	r.Run("localhost:8080")

}
