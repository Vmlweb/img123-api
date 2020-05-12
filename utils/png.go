package utils

import (
	"bytes"
	"image"
	"image/png"

	"github.com/pkg/errors"
)

// DecodePNG parse png format
func DecodePNG(b []byte) (image.Image, error) {
	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode png")
	}
	return img, nil
}

// EncodePNG parse png format
func EncodePNG(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, i); err != nil {
		return nil, errors.Wrap(err, "unable to encode png")
	}
	return buf.Bytes(), nil
}
