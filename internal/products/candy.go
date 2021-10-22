package products

import "time"

type candy struct {
	product
	flavour string
}

func NewCandy(flavour string) IProduct {
	return &candy{
		product: product{
			name:             "SUPER CANDY",
			command:          "GET-CANDY",
			price:            0.65,
			expirationDate:   time.Now().AddDate(4, 0, 0),
			weight:           15,
			extraInformation: "",
		},
		flavour: flavour,
	}
}

func (c *candy) ExtraInformation() string {
	return "Flavour: " + c.flavour
}
