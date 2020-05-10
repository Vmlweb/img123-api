package utils

import (
	"bytes"
	"image/jpeg"
	"image/png"

	"github.com/pkg/errors"
)

// JPGtoPNG JPG to PNG conversion
func JPGtoPNG(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, errors.Wrap(err, "unable to encode png")
	}

	return buf.Bytes(), nil
}
