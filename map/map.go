package maps

import "fmt"

func GenerateMapURL(latitude, longitude float64) string {
	const delta = 0.05
	MapURL := "https://www.openstreetmap.org/export/embed.html?bbox=%f,%f,%f,%f&layer=mapnik&marker=%f,%f"
	return fmt.Sprintf(
		MapURL,
		longitude-delta, latitude-delta, longitude+delta, latitude+delta, latitude, longitude,
	)
}
