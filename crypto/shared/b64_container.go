package shared

import (
	"encoding/base64"
	"strings"
)

// B64ContainerCreate creates a base64 encoded text container from a header and data.
//
// Parameters:
// - header: The header of the text container.
// - data: The data to be encoded.
// - lineLength: The length of the lines in the text container.
//
// Returns:
// - textContainer: The base64 encoded text container.
func B64ContainerCreate(header string, data []byte, lineLength int) (textContainer string) {
	header = strings.TrimSpace(header)
	b64String := base64.StdEncoding.EncodeToString(data)
	return TextContainerCreate(header, b64String, lineLength)
}

// B64ContainerParse parses a base64 encoded text container into a header and data.
//
// Parameters:
// - b64Container: The base64 encoded text container.
//
// Returns:
// - header: The header of the text container.
// - data: The decoded data.
// - err: An error if the base64 encoded text container is invalid.
func B64ContainerParse(b64Container string) (header string, data []byte, err error) {
	header, b64Body, err := TextContainerParse(b64Container)
	if err != nil {
		return "", nil, err
	}

	data, err = base64.StdEncoding.DecodeString(b64Body)
	if err != nil {
		return "", nil, err
	}

	return header, data, nil
}
