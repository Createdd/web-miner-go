package useCases

// import (
// 	InformationModel "github.com/Createdd/web-miner-go/entities"
// )

type fnString func(string) string

// FetchAndPresent takes an url and applies fetching and parsing function
func FetchAndPresent(url string, fetchFn fnString, parseFn fnString) string {
	/**
	 * Takes an url and applies it to a function for fetching.
	 * Then the fetched data is parsed and returned.
	 * TODO: Format the fetched data to fit the InformationModel
	 */
	data := fetchFn(url)
	parsedData := parseFn(data)

	return parsedData
}
