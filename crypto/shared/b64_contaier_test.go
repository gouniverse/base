package shared

import (
	"bytes"
	"testing"
)

func TestB64ContainerCreate(t *testing.T) {
	header := "Test Header"
	data := []byte("Test Data")
	lineLength := 76

	textContainer := B64ContainerCreate(header, data, lineLength)

	if textContainer == "" {
		t.Fatal("Failed to create text container")
	}

	if len(textContainer) != 25 {
		t.Fatal("Text container length is not 25, but: ", len(textContainer))
	}

	expected := "Test Header\n" + "VGVzdCBEYXRh\n"

	if textContainer != expected {
		t.Fatal("Text container does not match expected value recevewd:\n", textContainer, "\nexpected:\n", expected)
	}

	parsedHeader, parsedData, err := B64ContainerParse(textContainer)
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
