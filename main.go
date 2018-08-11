package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "works",
		})
	})
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"json": "works",
		})
	})
	router.Run(":3000")
	fmt.Println("_______________________")
	fmt.Println("hello")
}
