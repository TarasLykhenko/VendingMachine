package products

import "time"

type cola struct {
	product
}

func NewCola() IProduct {
	return &cola{
		product: product{
			name:             "Super Cola",
			command:          "GET-COLA",
			price:            1.35,
			expirationDate:   time.Now().AddDate(1, 0, 0),
			weight:           300,
			extraInformation: "Carbonation: 55%",
		},
	}
}
