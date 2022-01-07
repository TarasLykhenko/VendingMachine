package vendingmachine

import (
	"errors"
	"fmt"

	"github.com/TarasLykhenko/vending-machine-kata/internal/products"
)

//Coin .
type Coin interface {
	Mass() float32
	Diameter() float32
	Thickness() float32
}

type Currency struct {
	nickel  int
	dime    int
	quarter int
}

// VendingMachine .
type VendingMachine struct {
	status        Currency
	currentCredit float32
	cCoins        int
	bank          float32
	stock         []products.IProduct
	aproduct      float32
	change        float32
}

var cCoins int

// NewVendingMachine .
func NewVendingMachine(productList []products.IProduct) *VendingMachine {
	v := VendingMachine{}

	for _, product := range productList {
		v.stock = append(v.stock, product)
	}
	return &v
}

// AcceptCoin .
func (v *VendingMachine) AcceptCoin(coin Coin) {

	if !v.LimitReached(v.cCoins) {
		if coin.Mass() == 2.5 && coin.Diameter() == 19.05 && coin.Thickness() == 1.52 {
			fmt.Printf("I DO NOT ACCEPT PENNIES\n")
		}
		if coin.Mass() == 5 && coin.Diameter() == 21.21 && coin.Thickness() == 1.95 {
			v.currentCredit += 0.05
			v.status.nickel++
		}
		if coin.Mass() == 2.26 && coin.Diameter() == 17.91 && coin.Thickness() == 1.35 {
			v.currentCredit += 0.1
			v.status.dime++
		}
		if coin.Mass() == 6.25 && coin.Diameter() == 24.26 && coin.Thickness() == 1.75 {
			v.currentCredit += 0.25
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
	if v.currentCredit < product.Price() {
		fmt.Print("Price: " + fmt.Sprintf("%.2f", product.Price()) + "€\n")
		fmt.Print("INSERT COIN\n\n")
		return
	}

	err := v.CanMakeChange(product)

	if err != nil {
		fmt.Print("CANNOT BUY, EXACT CHANGE ONLY\n\n")
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
	v.currentCredit = v.currentCredit - productPrice
}

func (v *VendingMachine) CanMakeChange(product products.IProduct) error {
	_, err := v.CalculateChange(v.currentCredit - product.Price())

	if err != nil {
		return err
	}

	return nil
}

func (v *VendingMachine) CalculateChange(remainingCredit float32) (*Currency, error) {
	nickel := v.status.nickel
	dime := v.status.dime
	quarter := v.status.quarter

	nickelReturned := 0
	dimeReturned := 0
	quarterReturned := 0

	for remainingCredit >= 0.25 && quarter >= 1 {
		remainingCredit = remainingCredit - float32(0.25)
		quarter--
		quarterReturned++
	}
	for remainingCredit >= 0.10 && dime >= 1 {
		remainingCredit = remainingCredit - float32(0.10)
		dime--
		dimeReturned++
	}
	for remainingCredit >= 0.05 && nickel >= 1 {
		remainingCredit = remainingCredit - float32(0.05)
		nickel--
		nickelReturned++
	}

	if remainingCredit > 0.0 {
		return nil, errors.New("Not enough coins to return the change")
	}

	return &Currency{
		nickel:  nickel,
		dime:    dime,
		quarter: quarter,
	}, nil
}

// ReturnCoins .
func (v *VendingMachine) ReturnCoins() (*Currency, error) {
	coins, err := v.CalculateChange(v.currentCredit)

	if err != nil {
		return nil, err
	}

	v.currentCredit = 0
	v.status.nickel = v.status.nickel - coins.nickel
	v.status.dime = v.status.dime - coins.dime
	v.status.quarter = v.status.quarter - coins.quarter

	return coins, nil

}

// Display .
func (v *VendingMachine) Display() string {
	var output string

	output += fmt.Sprintf("Credit: %.2f€\n\n", v.currentCredit)

	hasChange := true
	for _, product := range v.stock {
		output += product.Name() + "\n"
		output += "||  " + product.ExtraInformation() + "\n"
		output += "||  Price: " + fmt.Sprintf("%.2f", product.Price()) + "€\n"
		output += "||  Command: " + product.Command()
		output += "\n"

		if hasChange == true {
			err := v.CanMakeChange(product)
			if err != nil {
				hasChange = false
			}
		}
	}

	output += "\n"
	if hasChange == true {
		output += "INSERT COIN\n\n"
	} else {
		output += "EXACT CHANGE ONLY\n\n"
	}

	return output

}

// RestockProducts .
func (v *VendingMachine) RestockProducts() {
}

// RetriveMoney .
func (v *VendingMachine) RetriveMoney() (*Currency, error) {

	nickelReturned := 0
	dimeReturned := 0
	quarterlReturned := 0

	for v.status.quarter > 4 {
		v.status.quarter--
		quarterlReturned++
	}
	for v.status.dime > 4 {
		v.status.dime--
		dimeReturned++
	}
	for v.status.nickel > 4 {
		v.status.nickel--
		nickelReturned++
	}
	retriveCoins := Currency{
		nickel:  nickelReturned,
		dime:    dimeReturned,
		quarter: quarterlReturned,
	}

	print("Quarters : ", quarterlReturned)
	print("\nDimers: ", dimeReturned)
	print("\nNickels: ", nickelReturned)

	return &retriveCoins, nil

}

func (v *VendingMachine) ReturnCommands() {
	fmt.Print("PENNY - Inserts a Penny\n")
	fmt.Print("NICKEL - Inserts a Nickel\n")
	fmt.Print("DIME - Inserts a Dime\n")
	fmt.Print("QUARTER - Inserts a Quarter\n")
	fmt.Print("GET-COLA - Get Cola\n")
	fmt.Print("GET-CHIPS - Get Chips\n")
	fmt.Print("GET-CANDY - Get Candy\n")
	fmt.Print("COIN-RETURN - Get your coins back\n")
	fmt.Print("RESTOCK - Vendor Restock of mnachine\n")
	fmt.Print("RETRIVE-MONEY - Vendor Retrieve coins\n")
	fmt.Print("HELP - This message\n\n")
}
