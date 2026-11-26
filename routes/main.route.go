package routes

import (
	"net/http"

	"meteo/controllers"
)

// MainRouter crée le routeur principal et y ajoute toutes les routes du site.
func MainRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Fichiers statiques : /static/css/style.css → assets/css/style.css
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	comuneRoutes(mux)

	// Toute autre URL : page 404 du site
	mux.HandleFunc("/", controllers.NotFound)
	return mux
}

func comuneRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/communes/suggestions", controllers.SuggestCommunes)
	mux.HandleFunc("GET /communes", controllers.ListCommunes)
	mux.HandleFunc("GET /communes/{id}", controllers.ShowCommune)
}