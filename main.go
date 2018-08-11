package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
	"io/ioutil"
	"strings"
	// "log"
)

func main() {

	// fmt.Println("_______________________")
	// fmt.Println("hello")

	resp, err := http.Get("https://medium.com/@ddcreationstudi/latest?format=json")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	data := string(body)

	convertedData := strings.Replace(data, "])}while(1);</x>", "", 1)

	fmt.Print(string(convertedData))
	fmt.Print("_________________")
	fmt.Print(err, err2)


	// Set up router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "works",
		})
	})
	router.GET("/json", func(c *gin.Context) {
		c.String(200, convertedData)
	})
	router.Run(":3000")

}
