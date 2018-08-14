package router

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func determineListenAddress() (string, error) {
	/**
	 * Get the Port from the environment or return an error.
	 * This is necessary for successful deplyoments.
	 * Returns a string or error.
	 */
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

// StartRouter creates the router
func StartRouter(data string) (err error) {
	/**
	 * StartRouter runs the router by determining the address to listen for from the PORT
	 * Set basic GET routes to run the app
	 */
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s...\n", addr)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "works",
		})
	})
	router.GET("/json", func(c *gin.Context) {
		c.String(200, data)
	})

	return router.Run(addr)
}
