package utils

import "github.com/teris-io/shortid"

// GenerateShortCode generates a unique short code for URLs
func GenerateShortCode() (string, error) {
	return shortid.Generate()
}
