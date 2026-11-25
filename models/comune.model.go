package models

import "errors"

var (
	ErrNotFound    = errors.New("commune introuvable")
	ErrBadRequest  = errors.New("requête invalide")
	ErrTimeout     = errors.New("délai dépassé")
	ErrUnavailable = errors.New("service indisponible")
)

type Commune struct {
	Name        string `json:"nom"`
	Code        string `json:"code"`
	Departement []Departement `json:"Departement"`
}

type Departement struct {
	Code string `json:"code"`
	Name  string `json:"nom"`
}