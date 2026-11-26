package main

import (
	"log"
	"net/http"

	"meteo/helpers"
	"meteo/routes"
)

func main() {
	helpers.LoadTemplates()
	router := routes.MainRouter()

	log.Println("Serveur lancé sur http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
