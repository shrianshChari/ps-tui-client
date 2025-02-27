package utils

import (
	"strings"
	"unicode"
)

// Removes all non-alphanumeric characters from a string and changes
// all the letter characters to lowercase
func ToID(s string) string {
	var sb strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	return strings.ToLower(sb.String())
}
