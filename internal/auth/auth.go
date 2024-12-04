package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey извлекает API ключ из заголовков запроса
// Пример:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.SplitN(val, " ", 2)
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
