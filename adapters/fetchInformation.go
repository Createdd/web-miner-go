package adapters

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FetchInformation shall fetch the information of a desired url
func FetchInformation(url string) (string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	data := string(body)
	return data
}