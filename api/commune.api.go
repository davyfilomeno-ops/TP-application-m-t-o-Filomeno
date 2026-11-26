package api

import (
	"net/url"

	"meteo/models"
)

const (
	baseURL   = "https://geo.api.gouv.fr"
	userAgent = "admin-geo"
)

// GetCommunesByName renvoie les communes correspondant au nom recherché.
func GetCommunesByName(nomCommune string) ([]models.Commune, error) {
	query := url.Values{}
	query.Set("nom", nomCommune)
	return getCommunes(query)
}

// GetCommune renvoie la commune correspondant à son code.
func GetCommune(code string) (models.Commune, error) {
	query := url.Values{}
	query.Set("code", code)

	communes, err := getCommunes(query)
	if err != nil {
		return models.Commune{}, err
	}
	if len(communes) == 0 {
		return models.Commune{}, models.ErrNotFound
	}
	return communes[0], nil
}

func getCommunes(query url.Values) ([]models.Commune, error) {
	query.Set("fields", "nom,code,centre")
	var communes []models.Commune
	err := get("/communes?"+query.Encode(), &communes)
	return communes, err
}
