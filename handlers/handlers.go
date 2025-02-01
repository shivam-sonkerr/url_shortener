package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
	"net/url"
	"url_shortener/dbutils"
	"url_shortener/models"
)

// Ping Handler
func HandlerPing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

// URL Redirect Handler
func RedirectURLHandler(c *gin.Context) {
	db, err := dbutils.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	shortURL := c.Param("shortURL")
	originalURL, err := dbutils.CheckIfURLExists(db, shortURL)

	if originalURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "shortURL has no valid mapping"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

// URL Post Handler
func URLPost(c *gin.Context) {
	var request models.URLMappings
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidURL(request.OriginalURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	db, err := dbutils.ConnectDB()
	if err != nil {
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate uniqueness of short URL"})
			return
		}
		if exists == 0 {
			break
		}
	}

	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Short URL created successfully",
		"shortURL": shortURL,
	})
	fmt.Println("Short URL: ", shortURL)
	c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("http://localhost:8080/redirect/%s", shortURL))
}

// Shorten and Redirect Handler
func ShortenAndRedirect(c *gin.Context) {
	var request models.URLMappings
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if request.OriginalURL == "" || !isValidURL(request.OriginalURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or empty URL"})
		return
	}

	db, err := dbutils.ConnectDB()
	if err != nil {
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		if exists == 0 {
			break
		}
	}

	_, err = dbutils.AddURL(db, request.OriginalURL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL mapping"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
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
