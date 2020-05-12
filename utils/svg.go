package utils

import (
	"bytes"
	"image"
	"image/png"

	b64 "encoding/base64"

	svg "github.com/ajstarks/svgo"
	"github.com/pkg/errors"
)

// EncodeSVG parse svg format
func EncodeSVG(i image.Image) ([]byte, error) {
	pngBuf := new(bytes.Buffer)
	err := png.Encode(pngBuf, i)
	if err != nil {
		return nil, errors.Wrap(err, "unable to encode svg")
	}

	base64 := b64.StdEncoding.EncodeToString(pngBuf.Bytes())

	config, _, err := image.DecodeConfig(pngBuf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to encode svg")
	}

	svgBuf := new(bytes.Buffer)
	svg := svg.New(svgBuf)
	svg.Start(config.Width, config.Height)
	svg.Image(0, 0, config.Width, config.Height, "data:image/png;base64,"+base64)
	svg.End()

	return svgBuf.Bytes(), nil
}
