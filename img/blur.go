package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

func Blur(content []byte, blur float64, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	// Create a blurred version of the image.
	dstImage := imaging.Blur(srcImage, blur)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}
