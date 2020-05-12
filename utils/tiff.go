package utils

import (
	"bytes"
	"image"

	"golang.org/x/image/tiff"

	"github.com/pkg/errors"
)

// DecodeTIFF decode tiff format
func DecodeTIFF(b []byte) (image.Image, error) {
	img, err := tiff.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode tiff")
	}
	return img, nil
}

// EncodeTIFF parse tiff format
func EncodeTIFF(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := tiff.Encode(buf, i, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode tiff")
	}
	return buf.Bytes(), nil
}
