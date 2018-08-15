package adapters

import (
	fetching "github.com/Createdd/web-miner-go/adapters/fetching"
	parsing "github.com/Createdd/web-miner-go/adapters/parsing"
	useCases "github.com/Createdd/web-miner-go/useCases"
)

// ProvideInformation shall convert data for further useage
func ProvideInformation() string {
	url := "https://medium.com/@ddcreationstudi/latest?format=json";

	// useCases.FetchAndPresent(url, "JSON")

	fetchFn := fetching.FetchInformation
	parseFn := parsing.ParseInformation

	convertedData := useCases.FetchAndPresent(url, fetchFn, parseFn)

	return convertedData
}
