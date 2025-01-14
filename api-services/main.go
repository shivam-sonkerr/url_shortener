package main

import (
	"github.com/gin-gonic/gin"
)

func handlerPing(c *gin.Context) {

	c.JSON(200, gin.H{"message": "pong"})
}

func main() {

	r := gin.Default()

	r.GET("/ping", handlerPing)

	r.Run("localhost:8080")

}
