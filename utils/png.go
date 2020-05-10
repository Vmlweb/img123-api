package utils

import (
	"bytes"
	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"github.com/chai2010/webp"

	"github.com/pkg/errors"
)

// PNGtoJPG PNG to JPG conversion
func PNGtoJPG(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode png")
	}

	return buf.Bytes(), nil
}

// PNGtoGIF PNG to GIF conversion
func PNGtoGIF(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode gif")
	}

	return buf.Bytes(), nil
}

// PNGtoBMP PNG to BMP conversion
func PNGtoBMP(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, img); err != nil {
		return nil, errors.Wrap(err, "unable to encode bmp")
	}

	return buf.Bytes(), nil
}

// PNGtoWEBP PNG to WEBP conversion
func PNGtoWEBP(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode webp")
	}

	return buf.Bytes(), nil
}

// PNGtoTIFF PNG to TIFF conversion
func PNGtoTIFF(imageBytes []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := tiff.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode tiff")
	}

	return buf.Bytes(), nil
}
