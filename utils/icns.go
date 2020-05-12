package utils

import (
	"bytes"
	"image"

	icns "github.com/jackmordaunt/icns"

	"github.com/pkg/errors"
)

// DecodeICNS parse icns format
func DecodeICNS(b []byte) (image.Image, error) {
	img, err := icns.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode icns")
	}
	return img, nil
}

// EncodeICNS parse icns format
func EncodeICNS(i image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := icns.Encode(buf, i); err != nil {
		return nil, errors.Wrap(err, "unable to encode icns")
	}
	return buf.Bytes(), nil
}
