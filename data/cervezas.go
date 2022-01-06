package data

import (
	"errors"

	"falabella.cl/prueba/model"
)

var (
	Cervezas = []model.Beer{
		model.Beer{
			Id:       1,
			Name:     "Golden",
			Brewery:  "Kross",
			Country:  "España",
			Price:    1.5,
			Currency: "EUR",
		},
		model.Beer{
			Id:       2,
			Name:     "Heineken",
			Brewery:  "CCU",
			Country:  "Chile",
			Price:    850,
			Currency: "CLP",
		},
		model.Beer{
			Id:       3,
			Name:     "Corona",
			Brewery:  "AB Inbev",
			Country:  "México",
			Price:    20.20,
			Currency: "MXN",
		},
	}
	ErrorBeer = errors.New("Error")
)
