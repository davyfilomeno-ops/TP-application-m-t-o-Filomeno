package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"

	"meteo/helpers"
	"meteo/models"
	"meteo/services"
)

func ListCommunes(w http.ResponseWriter, r *http.Request) {
	nomCommune := strings.TrimSpace(r.URL.Query().Get("nom"))

	communes, err := services.ListCommunes(nomCommune)
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", communes)
}

func SuggestCommunes(w http.ResponseWriter, r *http.Request) {
	prefix := strings.TrimSpace(r.URL.Query().Get("nom"))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if utf8.RuneCountInString(prefix) < 2 {
		if err := json.NewEncoder(w).Encode([]models.Commune{}); err != nil {
			log.Printf("Encodage des suggestions : %v", err)
		}
		return
	}

	communes, err := services.SuggestCommunes(r.Context(), prefix)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		log.Printf("Recherche de suggestions de communes : %v", err)
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, models.ErrBadRequest):
			status = http.StatusBadRequest
		case errors.Is(err, models.ErrTimeout), errors.Is(err, models.ErrUnavailable):
			status = http.StatusServiceUnavailable
		}
		w.WriteHeader(status)
		if encodeErr := json.NewEncoder(w).Encode(map[string]string{
			"error": "Les suggestions sont momentanément indisponibles.",
		}); encodeErr != nil {
			log.Printf("Encodage de l'erreur de suggestions : %v", encodeErr)
		}
		return
	}

	if communes == nil {
		communes = []models.Commune{}
	}
	if err := json.NewEncoder(w).Encode(communes); err != nil {
		log.Printf("Encodage des suggestions : %v", err)
	}
}

func ShowCommune(w http.ResponseWriter, r *http.Request) {
	communeMeteo, err := services.GetCommuneMeteo(r.PathValue("id"))
	if err != nil {
		renderErrorSevice(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "meteo", communeMeteo)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}


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
