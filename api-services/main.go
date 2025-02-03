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

	r.Static("/static", "./frontend")

	// Default route to serve index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	// URL shortening and redirection routes and also passed db as a variable
	r.POST("/shorten", func(c *gin.Context) { handlers.URLPost(c, db) })
	r.POST("/shorten-and-redirect", func(c *gin.Context) { handlers.ShortenAndRedirect(c, db) })

	// Route to handle the redirection of short URLs
	r.GET("/redirect/:shortURL", func(c *gin.Context) { handlers.RedirectURLHandler(c, db) })

	r.Run(":8080")
}
