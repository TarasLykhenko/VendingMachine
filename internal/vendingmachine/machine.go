package vendingmachine

import (
	"fmt"

	"github.com/TarasLykhenko/vending-machine-kata/internal/products"
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
	bank   float32
	stock  []products.IProduct
}

// NewVendingMachine .
func NewVendingMachine() *VendingMachine {
	v := VendingMachine{}
	v.startupStock()
	return &v
}

func (v *VendingMachine) startupStock() {
	v.stock = append(v.stock, products.NewCola())
	v.stock = append(v.stock, products.NewChips("Sea Salt"))
	v.stock = append(v.stock, products.NewChips("Bacon"))
	v.stock = append(v.stock, products.NewCandy("Blue"))
	v.stock = append(v.stock, products.NewCandy("Red"))
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
func (v *VendingMachine) SelectProduct(name string) {
	product := products.PrdFactory(name)
	if v.amoney <= product.Price() {
		fmt.Print("Price: " + fmt.Sprintf("%.2f", product.Price()) + "€\n")
		fmt.Print("INSERT COIN\n\n")
		return
	}

	v.MakeChange(product.Price())
	v.DispenseProduct(product)
	v.ReturnCoins()
	fmt.Print("Thank you!\n\n")
}

func (v *VendingMachine) DispenseProduct(p products.IProduct) {
	var index int
	for i, product := range v.stock {
		if product.Name() == p.Name() {
			index = i
			fmt.Print("BONK *NOISE OF " + product.Name() + " FALLING* \n")
			break
		}
	}

	v.stock[index] = v.stock[len(v.stock)-1] // Copy last element to index index.
	v.stock[len(v.stock)-1] = nil            // Erase last element (write zero value).
	v.stock = v.stock[:len(v.stock)-1]       // Truncate slice.
}

// MakeChange .
func (v *VendingMachine) MakeChange(productPrice float32) {
	v.bank += productPrice
	v.amoney = v.amoney - productPrice
}

// ReturnCoins .
func (v *VendingMachine) ReturnCoins() {
	if v.amoney != 0 {
		fmt.Print("CLINK CLINK *NOISE OF " + fmt.Sprintf("%.2f", v.amoney) + "€ IN COINS DROPING* \n")
		v.amoney = 0
	}
}

// Display .
func (v *VendingMachine) Display() string {
	var output string

	if v.amoney == 0 {
		output += "INSERT COIN\n"
	}

	output += fmt.Sprintf("Credit: %.2f€\n\n", v.amoney)

	for _, product := range v.stock {
		output += product.Name() + "\n"
		output += "||  " + product.ExtraInformation() + "\n"
		output += "||  Price: " + fmt.Sprintf("%.2f", product.Price()) + "€\n"
		output += "||  Command: " + product.Command()

		output += "\n\n"
	}

	return output

}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}

// RetriveMoney .
func (v *VendingMachine) RetriveMoney() {

}

func (v *VendingMachine) ReturnCommands() {
	fmt.Print("PENNY - Inserts a Penny\n")
	fmt.Print("NICKEL - Inserts a Nickel\n")
	fmt.Print("DIME - Inserts a Dime\n")
	fmt.Print("QUARTER - Inserts a Quarter\n")
	fmt.Print("GET-COLA - Get Cola\n")
	fmt.Print("GET-CHIPS - Get Chips\n")
	fmt.Print("GET-CANDY - Get Candy\n")
	fmt.Print("COIN RETURN - Get your coins back\n")
	fmt.Print("RESTOCK - Vendor Restock of mnachine\n")
	fmt.Print("RETRIVE-MONEY - Vendor Retrieve coins\n")
	fmt.Print("HELP - This message\n\n")
}
