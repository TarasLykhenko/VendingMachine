package products

import "time"

type chips struct {
	product
	flavour string
}

func NewChips(flavour string) IProduct {
	return &chips{
		product: product{
			name:             "SUPER CHIPS",
			command:          "GET-CHIPS",
			price:            0.5,
			expirationDate:   time.Now().AddDate(2, 0, 0),
			weight:           30,
			extraInformation: "% of air: 99%",
		},
		flavour: flavour,
	}
}

func (c *chips) ExtraInformation() string {
	return "Flavour: " + c.flavour + " | " + c.product.extraInformation
}
