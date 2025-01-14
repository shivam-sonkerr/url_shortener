package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type longURL struct {
	OriginalUrl string `json:"originalUrl"`
}

func handlerPing(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}

func urlPOST(c *gin.Context) {
	var originalUrl longURL

	if err := c.BindJSON(&originalUrl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Received URL:", originalUrl.OriginalUrl)

	return
}

func main() {

	r := gin.Default()

	r.GET("/ping", handlerPing)

	r.POST("/shorten", urlPOST)

	r.Run("localhost:8080")

}
