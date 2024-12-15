package qr

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gouniverse/base/img"
)

func UrlToQr(url string, width int, height int) []byte {
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)

	var buffer bytes.Buffer
	png.Encode(&buffer, qrCode)

	return buffer.Bytes()
}

func UrlToQrBase64(url string, width int, height int) string {
	qr := UrlToQr(url, width, height)
	qrBase64 := img.ToBase64Url(qr)
	return qrBase64
}
