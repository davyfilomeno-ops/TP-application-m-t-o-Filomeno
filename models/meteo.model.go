package models

type Meteo struct {
	Latitude             float64   `json:"latitude"`
	Longitude            float64   `json:"longitude"`
	GenerationTimeMS     float64   `json:"generationtime_ms"`
	UTCOffsetSeconds     int       `json:"utc_offset_seconds"`
	Timezone             string    `json:"timezone"`
	TimezoneAbbreviation string    `json:"timezone_abbreviation"`
	Elevation            float64   `json:"elevation"`
	DailyUnits           DailyUnits `json:"daily_units"`
	Daily                Daily      `json:"daily"`
}

type CommuneMeteo struct {
	Commune Commune `json:"commune"`
	Meteo   Meteo   `json:"meteo"`
}

type DailyUnits struct {
	Time             string `json:"time"`
	WeatherCode      string `json:"weather_code"`
	Temperature2mMin string `json:"temperature_2m_min"`
	Temperature2mMax string `json:"temperature_2m_max"`
}

type Daily struct {
	Time             []string  `json:"time"`
	WeatherCode      []int     `json:"weather_code"`
	Temperature2mMin []float64 `json:"temperature_2m_min"`
	Temperature2mMax []float64 `json:"temperature_2m_max"`
}