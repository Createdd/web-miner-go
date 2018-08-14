package router

import (
	"fmt"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)
func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

// StartRouter starts router
func StartRouter( data string) {

	fmt.Print("_________________")
	fmt.Print("works")

	addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
	}
  log.Printf("Listening on %s...\n", addr)

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
	router.Run(addr)
	// if err := router.Run((addr, nil); err != nil {
  //   panic(err)
  // }

}
