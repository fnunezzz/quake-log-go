package helpers

import "strings"

// Remove leading and trailing spaces from a string
func RemoveTrailingSpaces(text string) string {
	return strings.TrimSpace(text)

}