package shared

import (
	"encoding/hex"
	"strings"
)

// HexContainerCreate creates a text container from a header and data.
//
// Parameters:
// - header: The header of the text container.
// - data: The data to be hex encoded.
// - lineLength: The length of the lines in the text container.
//
// Returns:
// - textContainer: The text container.
func HexContainerCreate(header string, data []byte, lineLength int) (textContainer string) {
	header = strings.TrimSpace(header)
	hexString := hex.EncodeToString(data)
	return TextContainerCreate(header, hexString, lineLength)
}

// HexContainerParse parses a text container into a header and data.
//
// Parameters:
// - hexContainer: The hex text container.
//
// Returns:
// - header: The header of the text container.
// - data: The decoded data.
// - err: An error if the hex text container is invalid.
func HexContainerParse(hexContainer string) (header string, data []byte, err error) {
	header, hexBody, err := TextContainerParse(hexContainer)

	if err != nil {
		return "", nil, err
	}

	body, err := hex.DecodeString(hexBody)

	if err != nil {
		return "", nil, err
	}

	return header, body, nil
}
