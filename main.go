package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoseiji/go-weather/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handler.GetWeatherHandler).Methods(http.MethodGet)

	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", r)
}
