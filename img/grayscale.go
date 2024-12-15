package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

func Grayscale(content []byte, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	dstImage := imaging.Grayscale(srcImage)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}
