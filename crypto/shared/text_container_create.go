package shared

import (
	"strings"
)

// TextContainerCreate creates a text container from a header and body.
//
// Business Logic:
// - The header and body are trimmed of whitespace.
// - The header is the first line of the text container.
// - The body is the remaining lines of the text container.
// - The body is split into lines of the specified length.
// - The lines are joined with a newline character..
// - A single newline character is added to the end of the text container.
// - The text container is returned as a string.
//
// Note:
//   - The header and body are only trimmed of whitespace.
//   - The header and body are not encoded, any encoding must be done by the caller.
//   - The line length is the maximum length of the lines in the text container.
//   - The line length must be greater than 0.
//   - The line length does not have a maximum value, but 80 chars is recommended for readability.
//   - The new line at the end of the text container is added to keep the text container
//     consistent with the POSIX standard, stipulating that a text file is a sequence of
//     lines ending in a newline character.
//
// Parameters:
// - header: The header of the text container.
// - body: The body of the text container.
// - lineLength: The length of the lines in the text container.
//
// Returns:
// - The text container as a string.
func TextContainerCreate(header string, body string, lineLength int) string {
	header = strings.TrimSpace(header)
	body = strings.TrimSpace(body)

	lines := []string{
		header, // the header is the first line of the text container
	}

	for i := 0; i < len(body); i += lineLength {
		end := min(i+lineLength, len(body))
		lines = append(lines, body[i:end])
	}

	return strings.Join(lines, "\n") + "\n"
}
