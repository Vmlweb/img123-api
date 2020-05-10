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

// JPGtoGIF JPG to GIF conversion
func JPGtoGIF(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := gif.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode gif")
	}

	return buf.Bytes(), nil
}

// JPGtoBMP JPG to BMP conversion
func JPGtoBMP(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := bmp.Encode(buf, img); err != nil {
		return nil, errors.Wrap(err, "unable to encode bmp")
	}

	return buf.Bytes(), nil
}

// JPGtoWEBP JPG to WEBP conversion
func JPGtoWEBP(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode webp")
	}

	return buf.Bytes(), nil
}

// JPGtoTIFF JPG to TIFF conversion
func JPGtoTIFF(imageBytes []byte) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode jpeg")
	}

	buf := new(bytes.Buffer)
	if err := tiff.Encode(buf, img, nil); err != nil {
		return nil, errors.Wrap(err, "unable to encode tiff")
	}

	return buf.Bytes(), nil
}
