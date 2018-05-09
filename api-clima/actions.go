package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/briandowns/openweathermap"
	"github.com/gorilla/mux"
)

var OWM_API_KEY string = "8ea06e7900d3dc2de985e8d8a400fce0"

func Index(writerResponse http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tmpl.Execute(writerResponse, "Bienvenido a la API de Clima.")
	if err != nil {
		fmt.Fprintf(writerResponse, "No se pudo ejecutar la pagina de inicio.")
	}
}

func City(writerResponse http.ResponseWriter, request *http.Request) {

	parametros := mux.Vars(request)
	city := parametros["city"]

	tmpl, err := template.ParseFiles("templates/city.html")
	if err != nil {
		log.Fatalln(err)
	}

	w, err := openweathermap.NewCurrent("C", "ES", OWM_API_KEY)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(city)

	if w.Name == "" {
		fmt.Fprintf(writerResponse, "Ciudad invalida. Intente nuevamente.")
	} else {
		err = tmpl.Execute(writerResponse, w)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
