package vld

import "strings"

// isDate checks if a string is a valid date format
//
// The following formats are considered valid:
// - YYYY-MM-DD
//
// Business logic:
// - checks if the string contains 2 dashes
// - checks if the string does not contain colons
// - checks the first dash is at position 4
// - checks the second dash is at position 7
// - checks the string length is 10
func IsDate(value string) bool {
	countDashes := strings.Count(value, "-")

	if countDashes != 2 {
		return false
	}

	countColons := strings.Count(value, ":")

	if countColons > 0 {
		return false
	}

	if strings.Index(value, "-") != 4 {
		return false
	}

	if strings.LastIndex(value, "-") != 7 {
		return false
	}

	if len(value) != 10 {
		return false
	}

	return true
}
