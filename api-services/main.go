package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"url_shortener/dbutils"
	"url_shortener/handlers"
)

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

	r.Static("/static", "./frontend") // Serve static files under /static path

	// Default route to serve index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html") // Serve the index.html as default page
	})

	// URL shortening and redirection routes
	r.POST("/shorten", handlers.URLPost)
	r.POST("/shorten-and-redirect", handlers.ShortenAndRedirect)

	// Route to handle the redirection of short URLs
	r.GET("/:shortURL", handlers.RedirectURLHandler) // This will capture short URL and redirect

	r.Run(":8080")
}
