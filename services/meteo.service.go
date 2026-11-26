package services

import (
	"fmt"

	"meteo/api"
	"meteo/models"
)

// GetCommuneMeteo récupère la commune puis ses prévisions à son centre géographique.
func GetCommuneMeteo(code string) (models.CommuneMeteo, error) {
	commune, err := GetCommune(code)
	if err != nil {
		return models.CommuneMeteo{}, err
	}

	coordinates := commune.Centre.Coordinates
	if len(coordinates) < 2 {
		return models.CommuneMeteo{}, fmt.Errorf("coordonnées du centre absentes pour la commune %s", code)
	}

	weather, err := api.GetMeteo(coordinates[1], coordinates[0])
	if err != nil {
		return models.CommuneMeteo{}, err
	}

	return models.CommuneMeteo{
		Commune: commune,
		Meteo:   weather,
	}, nil
}