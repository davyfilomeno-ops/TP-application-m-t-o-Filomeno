package services

import (
	"context"
	"strings"

	"meteo/api"
	"meteo/models"
)

// ListCommunes renvoie les communes correspondant au nom recherché.
func ListCommunes(nomCommune string) ([]models.Commune, error) {
	return api.GetCommunesByName(nomCommune)
}

// SuggestCommunes conserve uniquement les communes dont le nom commence par le préfixe.
func SuggestCommunes(ctx context.Context, prefix string) ([]models.Commune, error) {
	communes, err := api.GetCommuneSuggestions(ctx, prefix)
	if err != nil {
		return nil, err
	}

	normalizedPrefix := normalizeCommuneName(prefix)
	suggestions := make([]models.Commune, 0, len(communes))
	for _, commune := range communes {
		if strings.HasPrefix(normalizeCommuneName(commune.Name), normalizedPrefix) {
			suggestions = append(suggestions, commune)
		}
	}
	return suggestions, nil
}

// GetCommune vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetCommune(code string) (models.Commune, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return models.Commune{}, models.ErrBadRequest
	}
	return api.GetCommune(code)
}

func normalizeCommuneName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	replacer := strings.NewReplacer(
		"à", "a", "á", "a", "â", "a", "ä", "a",
		"ç", "c",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"î", "i", "ï", "i",
		"ô", "o", "ö", "o",
		"ù", "u", "û", "u", "ü", "u",
		"ÿ", "y", "œ", "oe", "æ", "ae",
	)
	return replacer.Replace(name)
}