package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	utils "github.com/Vmlweb/img123-api/utils"
)

type supportedConversions struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// UtilitiesServer utilities server
type UtilitiesServer struct{}

// Supported supported
func (s *UtilitiesServer) Supported(w http.ResponseWriter, r *http.Request) {
	supported := []supportedConversions{
		{From: "jpg", To: "png"},
		{From: "jpg", To: "gif"},
		{From: "jpg", To: "bmp"},
		{From: "jpg", To: "webp"},
		{From: "jpg", To: "tiff"},
	}
	var jsonData []byte
	jsonData, err := json.Marshal(supported)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonData)
}

// Convert convert
func (s *UtilitiesServer) Convert(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var output []byte

	if from == "jpg" {
		if to == "png" {
			output, err = utils.JPGtoPNG(data)
		} else if to == "gif" {
			output, err = utils.JPGtoGIF(data)
		} else if to == "bmp" {
			output, err = utils.JPGtoBMP(data)
		} else if to == "webp" {
			output, err = utils.JPGtoWEBP(data)
		} else if to == "tiff" {
			output, err = utils.JPGtoTIFF(data)
		} else {
			http.Error(w, "incompatible conversion", 500)
		}
	} else if from == "png" {
		if to == "jpg" {
			output, err = utils.PNGtoJPG(data)
		} else if to == "gif" {
			output, err = utils.PNGtoGIF(data)
		} else if to == "bmp" {
			output, err = utils.PNGtoBMP(data)
		} else if to == "webp" {
			output, err = utils.PNGtoWEBP(data)
		} else if to == "tiff" {
			output, err = utils.PNGtoTIFF(data)
		} else {
			http.Error(w, "incompatible conversion", 500)
		}
	} else {
		http.Error(w, "incompatible conversion", 500)
	}

	if len(output) == 0 {
		http.Error(w, "incompatible conversion", 500)
	} else if err != nil {
		http.Error(w, err.Error(), 500)
	} else {
		w.Write(output)
	}
}
