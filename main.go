package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/test1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World Test1",
		})
	})

	router.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World Test2",
		})
	})

	router.Run(":3000")
}
