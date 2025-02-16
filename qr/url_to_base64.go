package qr

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gouniverse/base/img"
)

// UrlToQr returns a QR code for the given URL
//
// Parameters:
// - url string: the URL to generate the QR code for
// - width int: the width of the QR code
// - height int: the height of the QR code
//
// Returns:
// - []byte: the QR code as a byte slice
func UrlToQr(url string, width int, height int) []byte {
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)

	var buffer bytes.Buffer
	png.Encode(&buffer, qrCode)

	return buffer.Bytes()
}

// UrlToQrBase64 returns a Base64 encoded QR code for the given URL
//
// Parameters:
// - url string: the URL to generate the QR code for
// - width int: the width of the QR code
// - height int: the height of the QR code
//
// Returns:
// - string: the Base64 encoded QR code
func UrlToQrBase64(url string, width int, height int) string {
	qr := UrlToQr(url, width, height)
	qrBase64 := img.ToBase64Url(qr)
	return qrBase64
}
