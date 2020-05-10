package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

	utils "github.com/Vmlweb/img123-api/utils"
)

// UtilitiesServer utilities server
type UtilitiesServer struct{}

// Convert jpg to png converter
func (s *UtilitiesServer) Convert(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var output []byte

	if (from == "jpg" || from == "jpeg") && to == "png" {
		output, err = utils.JPGtoPNG(data)
		fmt.Println("JPG to PNG")
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
