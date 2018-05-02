package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
)

var OWM_API_KEY string = "8ea06e7900d3dc2de985e8d8a400fce0"

func Index(writerResponse http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writerResponse, "Bienvenido a la API de Clima.")
}

func Country(writerResponse http.ResponseWriter, request *http.Request) {
	// para obtener parametros del request
	parametros := mux.Vars(request)
	country := parametros["country"]
	state := parametros["state"]
	city := parametros["city"]

	fmt.Fprintln(writerResponse, "El pais es: "+country)
	fmt.Fprintln(writerResponse, "La provincia/estado es: "+state)
	fmt.Fprintln(writerResponse, "La ciudad es : "+city)

	log.Println(openweathermap.ValidAPIKey(OWM_API_KEY))

	w, err := openweathermap.NewCurrent("C", "ES", OWM_API_KEY)
	if err != nil {
		log.Fatalln(err)
	}

	//decoder := json.NewDecoder(w.CurrentByName("Phoenix,AZ"))

	w.CurrentByName("Phoenix")
	http.ur
	fmt.Fprintln(writerResponse, w)

}
