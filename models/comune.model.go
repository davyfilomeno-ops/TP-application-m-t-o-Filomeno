package models

type Commune struct {
	Name       string `json:"nom"`
	Code       string `json:"code"`
	Departement []Departement `json:"Departement"`
}

type Departement struct {
	Code string `json:"code"`
	Name  string `json:"nom"`
}