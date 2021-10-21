package vendingmachine

import (
	"fmt"
)

//Coin .
type Coin interface {
	Mass() float32
	Diameter() float32
	Thickness() float32
}

type Status struct {
	nickel  int
	dime    int
	quarter int
}

// VendingMachine .
type VendingMachine struct {
}

var money float32
var status Status

// NewVendingMachine .
func NewVendingMachine() *VendingMachine {
	return &VendingMachine{}
}

// AcceptCoin .
func (v *VendingMachine) AcceptCoin(coin Coin) {
	if coin.Mass() == 2.5 && coin.Diameter() == 19.05 && coin.Thickness() == 1.52 {
		fmt.Printf("I DO NOT ACCEPT PENNIES\n")
	}
	if coin.Mass() == 5 && coin.Diameter() == 21.21 && coin.Thickness() == 1.95 {
		money += 0.05
		status.nickel++
	}
	if coin.Mass() == 2.26 && coin.Diameter() == 17.91 && coin.Thickness() == 1.35 {
		money += 0.1
		status.dime++
	}
	if coin.Mass() == 6.25 && coin.Diameter() == 24.26 && coin.Thickness() == 1.75 {
		money += 0.25
		status.quarter++
	}
}

// SelectProduct .
func (v *VendingMachine) SelectProduct(product string) {

}

// MakeChange .
func (v *VendingMachine) MakeChange(money float32, productPrice float32) float32 {
	return money - productPrice
}

// ReturnCoins .
func (v *VendingMachine) ReturnCoins() {

}

// Display .
func (v *VendingMachine) Display() string {
	if money == 0 {
		return "INSERT COIN\n"
	}
	return fmt.Sprintf("CURRENT STATUS: $%f", money)
}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}

// RetriveMoney .
func (v *VendingMachine) RetriveMoney() {
}
