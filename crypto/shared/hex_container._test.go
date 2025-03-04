package shared

import (
	"bytes"
	"testing"
)

func TestHexContainerCreate(t *testing.T) {
	header := "Test Header"
	data := []byte("Test Data")
	lineLength := 76

	textContainer := HexContainerCreate(header, data, lineLength)

	if textContainer == "" {
		t.Fatal("Failed to create text container")
	}

	if len(textContainer) != 31 {
		t.Fatal("Text container length is not 31, but: ", len(textContainer))
	}

	expected := "Test Header\n" + "546573742044617461\n"

	if textContainer != expected {
		t.Fatal("Text container does not match expected value received:\n", textContainer, "\nexpected:\n", expected)
	}

	parsedHeader, parsedData, err := HexContainerParse(textContainer)
	if err != nil {
		t.Fatal("Failed to parse text container: ", err)
	}

	if parsedHeader != header {
		t.Fatal("Parsed header does not match original header")
	}

	if !bytes.Equal(parsedData, data) {
		t.Fatal("Parsed data does not match original data")
	}
}
