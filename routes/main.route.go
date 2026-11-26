package routes

import (
	"net/http"

	"meteo/controllers"
)

func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	comuneRoutes(mux)

	
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}

func comuneRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/communes/suggestions", controllers.SuggestCommunes)
	mux.HandleFunc("GET /communes", controllers.ListCommunes)
	mux.HandleFunc("GET /communes/{id}", controllers.ShowCommune)
}