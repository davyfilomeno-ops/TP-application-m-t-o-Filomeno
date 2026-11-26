package models

import "errors"

var (
	ErrNotFound    = errors.New("commune introuvable")
	ErrBadRequest  = errors.New("requête invalide")
	ErrTimeout     = errors.New("délai dépassé")
	ErrUnavailable = errors.New("service indisponible")
)

type Commune struct {
	Name   string `json:"nom"`
	Code   string `json:"code"`
	Centre Centre `json:"centre"`
}

type Centre struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}