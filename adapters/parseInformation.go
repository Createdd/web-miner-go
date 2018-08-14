package adapters

import "strings"

// ParseInformation shall convert data for further useage
func ParseInformation(data string) (string) {
	convertedData := strings.Replace(data, "])}while(1);</x>", "", 1)
	return convertedData
}