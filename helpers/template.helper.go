package helpers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates *template.Template

func LoadTemplates() {
	functions := template.FuncMap{
		"weatherIcon":  weatherIcon,
		"weatherLabel": weatherLabel,
		"formatDate":   formatDate,
	}
	templates = template.Must(template.New("").Funcs(functions).ParseGlob("templates/*.html"))
}

func weatherIcon(code int) string {
	switch {
	case code == 0:
		return "☀️"
	case code == 1:
		return "🌤️"
	case code == 2:
		return "⛅"
	case code == 3:
		return "☁️"
	case code == 45 || code == 48:
		return "🌫️"
	case code >= 51 && code <= 57:
		return "🌦️"
	case code >= 61 && code <= 67:
		return "🌧️"
	case code >= 71 && code <= 77:
		return "🌨️"
	case code >= 80 && code <= 82:
		return "🌦️"
	case code == 85 || code == 86:
		return "🌨️"
	case code >= 95 && code <= 99:
		return "⛈️"
	default:
		return "❓"
	}
}

func weatherLabel(code int) string {
	switch {
	case code == 0:
		return "Ciel dégagé"
	case code == 1:
		return "Majoritairement dégagé"
	case code == 2:
		return "Partiellement nuageux"
	case code == 3:
		return "Couvert"
	case code == 45 || code == 48:
		return "Brouillard"
	case code >= 51 && code <= 57:
		return "Bruine"
	case code >= 61 && code <= 67:
		return "Pluie"
	case code >= 71 && code <= 77:
		return "Neige"
	case code >= 80 && code <= 82:
		return "Averses"
	case code == 85 || code == 86:
		return "Averses de neige"
	case code >= 95 && code <= 99:
		return "Orages"
	default:
		return "Météo inconnue"
	}
}


func formatDate(value string) string {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		return value
	}

	days := [...]string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."}
	months := [...]string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."}
	return days[date.Weekday()] + " " + date.Format("2") + " " + months[date.Month()-1]
}

func Render(w http.ResponseWriter, status int, name string, data any) {
	var buf bytes.Buffer
	if err := templates.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("Erreur template (%s) : %v\n", name, err)
		RenderError(w, http.StatusInternalServerError, "Erreur chargement de la page")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buf.WriteTo(w)
}

type ErrorPage struct {
	Code    int
	Message string
}

func RenderError(w http.ResponseWriter, status int, message string) {
	Render(w, status, "error", ErrorPage{Code: status, Message: message})
}
