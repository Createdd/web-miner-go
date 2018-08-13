package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// StartRouter starts router
func StartRouter( data string) {
	// Set up router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "works",
		})
	})
	router.GET("/json", func(c *gin.Context) {
		c.String(200, data)
	})
	router.Run(":3000")

	fmt.Print("_________________")
	fmt.Print("works")

}
