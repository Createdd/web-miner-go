package parsing

import (
	"strings"
	// "github.com/Createdd/web-miner-go/useCases"
)

// ParseInformation shall convert data for further useage
func ParseInformation(data string) string {
	convertedData := strings.Replace(data, "])}while(1);</x>", "", 1)

	// useCases.PresentInformation
	return convertedData
}
