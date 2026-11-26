package api

import (
	"context"
	"net/url"

	"meteo/models"
)

const (
	baseURL   = "https://geo.api.gouv.fr"
	userAgent = "admin-geo"
)

func GetCommunesByName(nomCommune string) ([]models.Commune, error) {
	query := url.Values{}
	query.Set("nom", nomCommune)
	return getCommunes(query, "nom,code")
}

func GetCommuneSuggestions(ctx context.Context, prefix string) ([]models.Commune, error) {
	query := url.Values{}
	query.Set("nom", prefix)
	query.Set("limit", "10000")
	return getCommunesWithContext(ctx, query, "nom,code")
}

func GetCommune(code string) (models.Commune, error) {
	query := url.Values{}
	query.Set("code", code)

	communes, err := getCommunes(query, "nom,code,centre")
	if err != nil {
		return models.Commune{}, err
	}
	if len(communes) == 0 {
		return models.Commune{}, models.ErrNotFound
	}
	return communes[0], nil
}

func getCommunes(query url.Values, fields string) ([]models.Commune, error) {
	return getCommunesWithContext(context.Background(), query, fields)
}

func getCommunesWithContext(ctx context.Context, query url.Values, fields string) ([]models.Commune, error) {
	query.Set("fields", fields)
	var communes []models.Commune
	err := getWithContext(ctx, "/communes?"+query.Encode(), &communes)
	return communes, err
}
