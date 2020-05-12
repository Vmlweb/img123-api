package utils

import (
	"bytes"
	"image"

	"github.com/chai2010/webp"

	"github.com/pkg/errors"
)

// DecodeWEBP decode webp format
func DecodeWEBP(b []byte) (image.Image, error) {
	img, err := webp.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode webp")
	}
	return img, nil
}

// EncodeWEBP parse webp format
func EncodeWEBP(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, i, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode webp")
	}
	return buf.Bytes(), nil
}
