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
	status Status
	amoney float32
	cCoins int
}

// NewVendingMachine .
func NewVendingMachine() *VendingMachine {
	return &VendingMachine{}
}

// AcceptCoin .
func (v *VendingMachine) AcceptCoin(coin Coin) {

	if !v.LimitReached(v.cCoins) {
		if coin.Mass() == 2.5 && coin.Diameter() == 19.05 && coin.Thickness() == 1.52 {
			fmt.Printf("I DO NOT ACCEPT PENNIES\n")
		}
		if coin.Mass() == 5 && coin.Diameter() == 21.21 && coin.Thickness() == 1.95 {
			v.amoney += 0.05
			v.status.nickel++
		}
		if coin.Mass() == 2.26 && coin.Diameter() == 17.91 && coin.Thickness() == 1.35 {
			v.amoney += 0.1
			v.status.dime++
		}
		if coin.Mass() == 6.25 && coin.Diameter() == 24.26 && coin.Thickness() == 1.75 {
			v.amoney += 0.25
			v.status.quarter++
		}
	}
	v.cCoins++

}

func (v *VendingMachine) LimitReached(cCoins int) bool {
	if cCoins >= 100 {
		fmt.Printf("COINS LIMIT REACHED\n")
		return true
	}

	return false

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
	if v.amoney == 0 {
		return "INSERT COIN\n"
	}
	return fmt.Sprintf("CURRENT STATUS: $%f", v.amoney)
}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}

// RetriveMoney .
func (v *VendingMachine) RetriveMoney() {
}
