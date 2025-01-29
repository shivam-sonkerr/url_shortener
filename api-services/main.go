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

	r.GET("/ping", handlers.HandlerPing)

	r.POST("/shorten", handlers.URLPost)

	r.POST("/shorten-and-redirect", handlers.ShortenAndRedirect)

	r.Run(":8080")

}
