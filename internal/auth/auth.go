package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract auth token from authorization header and return
func GetApiKey(headers http.Header) (string, error) {
	head := headers.Get("Authorization")

	if head == "" {
		return "", errors.New("no authorization header was found")
	}

	data := strings.Split(head, " ")

	if len(data) != 2 {
		return "", errors.New("malformed authorization header. format like 'ApiKey --your-api-token--'")
	}

	if data[0] != "ApiKey" {
		return "", errors.New("malformed authorization header. format like 'ApiKey --your-api-token--'")
	}

	return data[1], nil
}
