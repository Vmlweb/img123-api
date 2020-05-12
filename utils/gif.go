package utils

import (
	"bytes"
	"image"
	"image/gif"

	"github.com/pkg/errors"
)

// DecodeGIF decode gif format
func DecodeGIF(b []byte) (image.Image, error) {
	img, err := gif.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode gif")
	}
	return img, nil
}

// EncodeGIF parse gif format
func EncodeGIF(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, i, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode gif")
	}
	return buf.Bytes(), nil
}
