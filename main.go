package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	app "github.com/Vmlweb/img123-api/app"
)

func main() {
	app := app.UtilitiesServer{}
	r := mux.NewRouter()

	r.HandleFunc("/convert", app.Convert).Methods("POST")
	r.HandleFunc("/supported", app.Supported).Methods("GET")

	handler := cors.Default().Handler(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on 0.0.0.0:" + port)
	error := http.ListenAndServe("0.0.0.0:"+port, handler)
	if error != nil {
		fmt.Print(error)
	}
}
