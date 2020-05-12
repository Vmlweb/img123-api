package utils

import (
	"bytes"
	"image"

	"golang.org/x/image/bmp"

	"github.com/pkg/errors"
)

// DecodeBMP parse bmp format
func DecodeBMP(b []byte) (image.Image, error) {
	img, err := bmp.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode bmp")
	}
	return img, nil
}

// EncodeBMP parse bmp format
func EncodeBMP(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, i); err != nil {
		return nil, errors.Wrap(err, "unable to encode bmp")
	}
	return buf.Bytes(), nil
}
