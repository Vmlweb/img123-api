package utils

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/pkg/errors"
)

// DecodeJPG decode jpeg format
func DecodeJPG(b []byte) (image.Image, error) {
	img, err := jpeg.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpg")
	}
	return img, nil
}

// EncodeJPG parse jpeg format
func EncodeJPG(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, i, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode jpg")
	}
	return buf.Bytes(), nil
}
