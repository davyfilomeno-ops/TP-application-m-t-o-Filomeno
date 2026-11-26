package services

import (
	"strings"

	"meteo/api"
	"meteo/models"
)

// ListCommunes renvoie les communes correspondant au nom recherché.
func ListCommunes(nomCommune string) ([]models.Commune, error) {
	return api.GetCommunesByName(nomCommune)
}

// GetCommune vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetCommune(code string) (models.Commune, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(code)
}