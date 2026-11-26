package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"meteo/helpers"
	"meteo/models"
	"meteo/services"
)

// ListCommunes affiche la liste des communes de la première génération.
func ListCommunes(w http.ResponseWriter, r *http.Request) {
	nomCommune := strings.TrimSpace(r.URL.Query().Get("nom"))
	if nomCommune == "" {
		nomCommune = ""
	}

	communes, err := services.ListCommunes(nomCommune)
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", communes)
}

// ShowCommune affiche la fiche de la commune dont l'identifiant est dans l'URL.
func ShowCommune(w http.ResponseWriter, r *http.Request) {
	communeMeteo, err := services.GetCommuneMeteo(r.PathValue("id"))
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "meteo", communeMeteo)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderErrorSevice(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound, "Cette commune n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest, "Numéro de commune invalide.")
	case errors.Is(err, models.ErrTimeout), errors.Is(err, models.ErrUnavailable):
		log.Println("API externe :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable, "Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError, "Une erreur est survenue.")
	}
}
