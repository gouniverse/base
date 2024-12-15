package img

import (
	"log"
	"os"
)

// ImgToBase64Url converts an image file to Base64 encoded URL string
func ImgToBase64Url(filePath string) string {
	// Read the entire file into a byte slice
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return ""
	}

	return ToBase64Url(bytes)
}
