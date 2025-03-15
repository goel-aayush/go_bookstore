package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParseBody decodes the request body into the provided struct (interface{}).
// It returns an error if parsing fails.
func ParseBody(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(x)
	if err != nil {
		return fmt.Errorf("failed to parse request body: %v", err)
	}
	return nil
}
