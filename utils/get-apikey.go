package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	authValues := strings.Split(authorization, " ")
	if authValues[0] == "" {
		return "", errors.New("no auth found")
	}

	if authValues[0] != "ApiKey" {
		return "", errors.New("malformed auth")
	}

	return authValues[1], nil

}
