package api

import (
	"net/url"
	"strconv"

	"meteo/models"
)

const meteoBaseURL = "https://api.open-meteo.com/v1/forecast"

func GetMeteo(latitude, longitude float64) (models.Meteo, error) {
	query := url.Values{}
	query.Set("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	query.Set("longitude", strconv.FormatFloat(longitude, 'f', -1, 64))
	query.Set("daily", "weather_code,temperature_2m_min,temperature_2m_max")
	query.Set("timezone", "auto")
	query.Set("forecast_days", "7")

	var meteo models.Meteo
	err := getURL(meteoBaseURL+"?"+query.Encode(), &meteo)
	return meteo, err
}