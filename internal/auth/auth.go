package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extract an API key from the headers if an HTTP request
// Example:
// Authorization: ApiKey {insert api key here}
func GetApiKey(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")
	if authorizationHeader == "" {
		return "", errors.New("auth header not found")
	}

	values := strings.Split(authorizationHeader, " ")
	if len(values) != 2 {
		return "", errors.New("incorrect format for auth header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("incorrect format for auth header - first part")
	}

	return values[1], nil
}
