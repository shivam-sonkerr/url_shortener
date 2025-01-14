package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
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

func urlGenerator() {
	random := rand.Int()

	prefix := "simple-url"
	fmt.Println(random)

	concat := fmt.Sprintf("%s%d", prefix, random)

	fmt.Println(concat)
}

func main() {

	r := gin.Default()

	r.GET("/ping", handlerPing)

	r.POST("/shorten", urlPOST)

	r.Run("localhost:8080")

}
