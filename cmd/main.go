package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	err := router.Run()
	if err != nil {
		log.Fatal("failed to start HTTP server")
		return
	}
}
