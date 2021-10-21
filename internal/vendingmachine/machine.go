package vendingmachine

import (
	"fmt"
	"strconv"

	"github.com/TarasLykhenko/vending-machine-kata/internal/products"
)

//Coin .
type Coin interface {
	Mass() float32
	Diameter() float32
	Thickness() float32
}

// VendingMachine .
type VendingMachine struct {
}

// NewVendingMachine .
func NewVendingMachine() *VendingMachine {
	return &VendingMachine{}
}

// AcceptCoin .
func (v *VendingMachine) AcceptCoin(coin Coin) {

}

// SelectProduct .
func (v *VendingMachine) SelectProduct(product string) {

}

// MakeChange .
func (v *VendingMachine) MakeChange() {

}

// ReturnCoins .
func (v *VendingMachine) ReturnCoins() {

}

// Display .
func (v *VendingMachine) Display() string {
	var output string
	list := []products.IProduct{}

	list = append(list, products.NewCola())
	list = append(list, products.NewChips("Sea Salt"))
	list = append(list, products.NewChips("Bacon"))
	list = append(list, products.NewCandy("Blue"))
	list = append(list, products.NewCandy("Red"))

	output = "INSERT COIN:\n"
	for i, product := range list {
		output += strconv.Itoa(i) + ". " + product.Name() + "\n"
		output += "  " + product.ExtraInformation() + "\n"
		output += "  Price: " + fmt.Sprintf("%.2f", product.Price()) + "€\n"

		output += "\n"
	}

	return output
}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}

// RetriveMoney .
func (v *VendingMachine) RetriveMoney() {
}
