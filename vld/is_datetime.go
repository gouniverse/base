package vld

import "strings"

// isDateTime checks if a string is a valid datetime format
//
// The following formats are condsidered valid:
// - YYYY-MM-DD HH:MM:SS
// - YYYY-MM-DDTHH:MM:SS
// - YYYY-MM-DDTHH:MM:SSZ
//
// Business logic:
// - checks if the string contains 2 dashes
// - checks if the string contains 2 colons
// - checks the first dash is at position 4
// - checks the second dash is at position 7
// - checks the first colon is at position 10
// - checks the second colon is at position 13
// - checks the string length is 19 or 20
// - checks the last character is 'Z' if the string length is 20
func IsDateTime(value string) bool {
	countDashes := strings.Count(value, "-")

	if countDashes != 2 {
		return false
	}

	countColons := strings.Count(value, ":")

	if countColons != 2 {
		return false
	}

	if strings.Index(value, "-") != 4 {
		return false
	}

	if strings.LastIndex(value, "-") != 7 {
		return false
	}

	if strings.Index(value, ":") != 13 {
		return false
	}

	if strings.LastIndex(value, ":") != 16 {
		return false
	}

	if len(value) != 19 && len(value) != 20 {
		return false
	}

	if len(value) == 20 && value[19] != 'Z' {
		return false
	}

	return true
}
