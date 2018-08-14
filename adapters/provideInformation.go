package adapters

import (
	fetching "github.com/Createdd/web-miner-go/adapters/fetching"
	parsing "github.com/Createdd/web-miner-go/adapters/parsing"
	// "github.com/Createdd/web-miner-go/useCases"
)

// ProvideInformation shall convert data for further useage
func ProvideInformation() string {

	data := fetching.FetchInformation("https://medium.com/@ddcreationstudi/latest?format=json")
	convertedData := parsing.ParseInformation(data)

	return convertedData
}
