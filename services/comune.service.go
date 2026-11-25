import (
	"strconv"

	"meteo/api"
	"meteo/models"
)

// GetFirstGeneration renvoie les communes de la première génération.
func GetFirstGeneration() ([]models.Comune, error) {
	return api.GetGeneration(1)
}

// GetComune vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetComune(idParam string) (models.Comune, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Comune{}, models.ErrBadRequest
	}
	return api.GetComune(id)
}
