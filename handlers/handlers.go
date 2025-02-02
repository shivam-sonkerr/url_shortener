package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"url_shortener/dbutils"
	"url_shortener/models"
)

// RedirectURLHandler URL Redirect Handler
func RedirectURLHandler(c *gin.Context) {
	log.Println("RedirectURLHandler called")

	db, err := dbutils.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	shortURL := c.Param("shortURL")
	log.Println("Received short URL:", shortURL)

	originalURL, err := dbutils.CheckIfURLExists(db, shortURL)
	if err != nil {
		log.Println("Database query error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}

	if originalURL == "" {
		log.Println("Short URL not found:", shortURL)
		c.JSON(http.StatusNotFound, gin.H{"error": "shortURL has no valid mapping"})
		return
	}

	log.Println("Redirecting to:", originalURL)
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

// URL Post Handler
func URLPost(c *gin.Context) {
	log.Println("URLPost handler called")

	var request models.URLMappings
	if err := c.BindJSON(&request); err != nil {
		log.Println("Invalid request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Parsed Request:", request)

	if !isValidURL(request.OriginalURL) {
		log.Println("Invalid URL format:", request.OriginalURL)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	db, err := dbutils.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database."})
		return
	}
	defer db.Close()

	var shortURL string
	for {
		shortURL = urlGenerator()
		var exists int
		err = db.QueryRow("SELECT COUNT(*) FROM url_mappings where short_url = ?", shortURL).Scan(&exists)
		if err != nil {
			log.Println("Error checking short URL uniqueness:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate uniqueness of short URL"})
			return
		}
		if exists == 0 {
			break
		}
	}

	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		log.Println("Failed to save URL mapping:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	log.Println("Short URL created:", shortURL)
	c.JSON(http.StatusCreated, gin.H{
		"message":       "Short URL created successfully",
		"shortened_url": fmt.Sprintf("http://localhost:8080/redirect/%s", shortURL),
	})
}

// Shorten and Redirect Handler
func ShortenAndRedirect(c *gin.Context) {
	log.Println("ShortenAndRedirect handler called")

	var request models.URLMappings
	if err := c.BindJSON(&request); err != nil {
		log.Println("Invalid request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if request.OriginalURL == "" || !isValidURL(request.OriginalURL) {
		log.Println("Invalid or empty URL provided:", request.OriginalURL)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or empty URL"})
		return
	}

	db, err := dbutils.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	var shortURL string
	for {
		shortURL = urlGenerator()
		var exists int
		err = db.QueryRow("SELECT COUNT(*) FROM url_mappings WHERE short_url = ?", shortURL).Scan(&exists)
		if err != nil {
			log.Println("Database error while checking uniqueness:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		if exists == 0 {
			break
		}
	}

	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		log.Println("Failed to save URL mapping:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	log.Println("Shortened URL generated:", shortURL)
	c.JSON(http.StatusOK, gin.H{
		"shortened_url": fmt.Sprintf("http://localhost:8080/redirect/%s", shortURL),
	})
}

// Helper Functions
func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func urlGenerator() string {
	random := rand.IntN(9000) + 1000
	prefix := "simple-url"
	shortURL := fmt.Sprintf("%s-%d", prefix, random)
	return shortURL
}
