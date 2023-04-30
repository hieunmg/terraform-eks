package utils

import "strings"

func ExtractTokenSignature(token string) string {
	parts := strings.Split(token, ".")
	return parts[2]
}
