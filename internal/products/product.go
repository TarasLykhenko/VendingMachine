package products

import "time"

type IProduct interface {
	Name() string
	Command() string
	Price() float32
	ExpirationDate() time.Time
	Weight() float32
	ExtraInformation() string
}

type product struct {
	name             string
	command          string
	price            float32
	expirationDate   time.Time
	weight           float32
	extraInformation string
}

// NewProduct creates a new product
func NewProduct(name string, command string, price float32, expirationDate time.Time, weight float32, extraInformation string) *product {
	return &product{
		name:             name,
		command:          command,
		price:            price,
		expirationDate:   expirationDate,
		weight:           weight,
		extraInformation: extraInformation,
	}
}

func PrdFactory(name string) IProduct {
	switch name {
	case "COLA":
		return NewCola()
	case "CHIPS":
		return NewChips("")
	case "CANDY":
		return NewCandy("")
	default:
		return nil
	}
}

// getters
func (p *product) Name() string {
	return p.name
}
func (p *product) Command() string {
	return p.command
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
