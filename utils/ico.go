package utils

import (
	"bytes"
	"image"

	ico "github.com/keybase/golang-ico"

	"github.com/pkg/errors"
)

// DecodeICO parse ico format
func DecodeICO(b []byte) (image.Image, error) {
	img, err := ico.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode ico")
	}
	return img, nil
}

// EncodeICO parse ico format
func EncodeICO(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ico.Encode(buf, i); err != nil {
		return nil, errors.Wrap(err, "unable to encode ico")
	}
	return buf.Bytes(), nil
}
