package products

import "time"

type IProduct interface {
	Name() string
	Price() float32
	ExpirationDate() time.Time
	Weight() float32
	ExtraInformation() string
}

type product struct {
	name             string
	price            float32
	expirationDate   time.Time
	weight           float32
	extraInformation string
}

// NewProduct creates a new product
func NewProduct(name string, price float32, expirationDate time.Time, weight float32, extraInformation string) *product {
	return &product{
		name:             name,
		price:            price,
		expirationDate:   expirationDate,
		weight:           weight,
		extraInformation: extraInformation,
	}
}

// getters
func (p *product) Name() string {
	return p.name
}
func (p *product) Price() float32 {
	return p.price
}
func (p *product) ExpirationDate() time.Time {
	return p.expirationDate
}
func (p *product) Weight() float32 {
	return p.weight
}
func (p *product) ExtraInformation() string {
	return p.extraInformation
}
