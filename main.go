package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/Createdd/web-miner-go/router"
)

func main() {
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

	router.StartRouter(convertedData)

	fmt.Print(string(convertedData))
}
