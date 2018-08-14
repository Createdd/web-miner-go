package router

import (
	"fmt"
	"log"
	"os"

	"github.com/Createdd/web-miner-go/adapters"
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
func StartRouter() (err error) {
	/**
	 * StartRouter runs the router by determining the address to listen for from the PORT
	 * Set basic GET routes to run the app
	 */
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s...\n", addr)

	data := adapters.FetchInformation("https://medium.com/@ddcreationstudi/latest?format=json")
	convertedDate := adapters.ParseInformation(data)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "works",
		})
	})

	router.GET("/json", func(c *gin.Context) {
		c.String(200, convertedDate)
	})

	return router.Run(addr)
}
