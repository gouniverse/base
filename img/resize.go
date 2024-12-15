package img

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

func Resize(content []byte, width, height int, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}
