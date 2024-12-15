package img

import (
	"encoding/base64"
	"net/http"
)

// ToBase64Url converts a byte slice to Base64 encoded URL string
func ToBase64Url(imgBytes []byte) string {
	mimeType := http.DetectContentType(imgBytes)

	base64Encoding := ""

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/bmp":
		base64Encoding += "data:image/bmp;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}

	base64Encoding += base64Encode(imgBytes)
	return base64Encoding
}

func base64Encode(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
}
