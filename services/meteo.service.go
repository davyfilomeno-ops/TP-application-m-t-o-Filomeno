package services

import (
	"fmt"

	"meteo/api"
	"meteo/map"
	"meteo/models"
)

func GetCommuneMeteo(code string) (models.PageData, error) {
	commune, err := GetCommune(code)
	if err != nil {
		return models.PageData{}, err
	}

	coordinates := commune.Centre.Coordinates
	if len(coordinates) < 2 {
		return models.PageData{}, fmt.Errorf("coordonnées du centre absentes pour la commune %s", code)
	}

	weather, err := api.GetMeteo(coordinates[1], coordinates[0])
	if err != nil {
		return models.PageData{}, err
	}

	fmt.Println(maps.GenerateMapURL(coordinates[1], coordinates[0]))

	return models.PageData{
		Commune: commune,
		Meteo:   weather,
		MapURL:  maps.GenerateMapURL(coordinates[1], coordinates[0]),
	}, nil
}
