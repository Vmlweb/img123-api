package app

import (
	"encoding/json"
	"image"
	"io/ioutil"
	"log"
	"net/http"

	utils "github.com/Vmlweb/img123-api/utils"
)

type converter struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type converterMethods struct {
	name   string
	decode func(d []byte) (image.Image, error)
	encode func(i image.Image) ([]byte, error)
}

var converters = []*converterMethods{
	{
		name:   "jpg",
		decode: utils.DecodeJPG,
		encode: utils.EncodeJPG,
	},
	{
		name:   "png",
		decode: utils.DecodePNG,
		encode: utils.EncodePNG,
	},
	{
		name:   "bmp",
		decode: utils.DecodeBMP,
		encode: utils.EncodeBMP,
	},
	{
		name:   "gif",
		decode: utils.DecodeGIF,
		encode: utils.EncodeGIF,
	},
	{
		name:   "tiff",
		decode: utils.DecodeTIFF,
		encode: utils.EncodeTIFF,
	},
	{
		name:   "webp",
		decode: utils.DecodeWEBP,
		encode: utils.EncodeWEBP,
	},
	{
		name:   "ico",
		decode: utils.DecodeICO,
		encode: utils.EncodeICO,
	},
	{
		name:   "icns",
		decode: utils.DecodeICNS,
		encode: utils.EncodeICNS,
	},
	{
		name:   "svg",
		encode: utils.EncodeSVG,
	},
}

// UtilitiesServer utilities server
type UtilitiesServer struct{}

// Supported generate map of supported conversions
func (s *UtilitiesServer) Supported(w http.ResponseWriter, r *http.Request) {
	converterMap := &[]converter{}
	for _, from := range converters {
		for _, to := range converters {
			if from.name != to.name && from.decode != nil && to.encode != nil {
				*converterMap = append(*converterMap, converter{
					From: from.name,
					To:   to.name,
				})
			}
		}
	}

	var jsonData []byte
	jsonData, err := json.Marshal(converterMap)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonData)
}

// Convert image
func (s *UtilitiesServer) Convert(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	var decoder *converterMethods
	var encoder *converterMethods
	for _, converter := range converters {
		if converter.name == from && converter.decode != nil {
			decoder = converter
		}
		if converter.name == to && converter.encode != nil {
			encoder = converter
		}
	}

	if decoder == nil || encoder == nil {
		http.Error(w, "incompatible conversion", 500)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	image, err := decoder.decode(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := encoder.encode(image)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if len(output) == 0 {
		http.Error(w, "incompatible conversion", 500)
	} else {
		w.Write(output)
	}
}
