package services

import (
	"strconv"

	"meteo/api"
	"meteo/models"
)

// GetFirstGeneration renvoie les communes de la première génération.
func GetFirstGeneration() ([]models.Commune, error) {
	return api.GetGeneration(1)
}

// GetCommune vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetCommune(idParam string) (models.Commune, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(id)
}
