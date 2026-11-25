package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"meteo/models"
)

// URL de base de l'API et identification de notre programme.
const (
	baseURL   = "https:https://geo.api.gouv.fr/decoupage-administratif/communes#name"
	userAgent = "admin-geo"
)

// Un seul client HTTP, partagé par toutes les requêtes.
var client = &http.Client{
	Timeout: 5 * time.Second,
}

// get envoie une requête GET vers l'endpoint et décode la réponse dans dest.
func get(endpoint string, dest any) error {
	// 1. Préparer la requête : méthode, URL et corps (nil pour un GET).
	request, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("création de la requête : %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)

	// 2. Envoyer la requête avec le client.
	response, err := client.Do(request)
	if err != nil {
		// Aucune réponse reçue : délai dépassé, réseau, DNS…
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return models.ErrTimeout
		}
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()

	// 3. Vérifier le code de statut avant de lire le corps.
	switch {
	case response.StatusCode == http.StatusOK:
		// 200 : succès, on peut décoder le corps.
	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode == http.StatusBadRequest:
		return models.ErrBadRequest
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	// 4. Décoder le flux JSON directement dans la variable de destination.
	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}

// GetGeneration renvoie les communes de la première génération (endpoint /communes).
func GetGeneration(generation int) ([]models.Comune, error) {
	var communes []models.Comune
	err := get("/communes", &communes)
	return communes, err
}

// GetComune renvoie la fiche d'une commune (endpoint /communes/{id}).
func GetComune(id int) (models.Comune, error) {
	var comune models.Comune
	err := get(fmt.Sprintf("/communes/%d", id), &comune)
	return comune, err
}