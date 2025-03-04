package shared

import (
	"errors"
	"strings"
	"unicode"
)

// TextContainerParse parses the input text container into a header and body.
//
// Business logic:
// - the text container must have a header and body.
// - the header and body must be separated by a newline.
// - the body may contain newlines, for keeping the readability of the body.
// - all newlines and carriage returns must be removed from the body before decoding.
//
// Parameters:
//
//   - textContainer: The input string containing the header and body.
//
// Returns:
//
//   - header: The parsed header string.
//   - body: The parsed body string.
//   - err: An error if the input format is invalid.
func TextContainerParse(textContainer string) (string, string, error) {
	// Trim leading/trailing whitespace immediately.
	//textContainer = strings.ReplaceAll(textContainer, "\n", "NEWLINE")
	//textContainer = sanitize(textContainer)
	//textContainer = strings.ReplaceAll(textContainer, "NEWLINE", "\n")
	textContainer = strings.TrimSpace(textContainer)

	// Handle empty input
	if textContainer == "" {
		return "", "", errors.New("empty text container")
	}

	// Split into header and remaining parts.
	parts := strings.Split(textContainer, "\n")

	header := parts[0] // Always at least one part exists

	body := ""
	if len(parts) > 1 {
		// Combine all body lines, then sanitize.
		body = strings.Join(parts[1:], "")
	}

	header = strings.TrimSpace(header)

	if header == "" {
		return "", "", errors.New("empty header")
	}

	if hasInvalidChars(header) {
		return "", "", errors.New("invalid characters in header")
	}

	body = strings.ReplaceAll(body, " ", "")
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\r", "")
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.TrimSpace(body)

	if hasInvalidChars(body) {
		return "", "", errors.New("invalid characters in body")
	}

	return header, body, nil
}

// hasInvalidChars checks if a string contains non-letter, non-digit, and non-printable characters.
func hasInvalidChars(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && (r < 32 || r > 126) {
			return true
		}
	}

	return false
}

// sanitize removes non-letter, non-digit, and non-printable characters from a string.
func sanitize(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || (r >= 32 && r <= 126) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
