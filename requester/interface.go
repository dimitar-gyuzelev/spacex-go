package requester

import (
	"fmt"
	"strings"
)

type Result interface {}

type Payload interface {}

type Headers map[string]string

type QueryStrings map[string]string
func (qs QueryStrings) String() string {
	var sb strings.Builder
	for k, v := range qs {
		sb.WriteString(fmt.Sprintf("%s=%s&", k, v))
	}
	return strings.TrimRight(sb.String(), "&")
}

type Config struct {
	URL string
	Headers Headers
	QueryStrings QueryStrings
}

type Requester interface {
	// Get performs a GET request.
	// Response is stored in result.
	Get(config Config, result Result) error

	// Post performs a POST request.
	// Response is store in result
	Post(config Config, payload Payload, result Result) error
}