package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TarasLykhenko/vending-machine-kata/internal/coins"
	"github.com/TarasLykhenko/vending-machine-kata/internal/products"
	"github.com/TarasLykhenko/vending-machine-kata/internal/vendingmachine"
)

func main() {

	productList := []products.IProduct{}

	productList = append(productList, products.NewCola())
	productList = append(productList, products.NewChips("Sea Salt"))
	productList = append(productList, products.NewChips("Bacon"))
	productList = append(productList, products.NewCandy("Blue"))
	productList = append(productList, products.NewCandy("Red"))

	vm := vendingmachine.NewVendingMachine(productList)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(vm.Display())
		fmt.Print("Input: ")
		text, _ := reader.ReadString('\n')
		trimtext := strings.Trim(text, "\n")
		fmt.Println(text)
		commands := strings.Split(strings.ToUpper(trimtext), ",")

		for _, c := range commands {
			switch c {
			case "PENNY":
				vm.AcceptCoin(coins.NewCoin(2.5, 19.05, 1.52))
			case "NICKEL":
				vm.AcceptCoin(coins.NewCoin(5, 21.21, 1.95))
			case "DIME":
				vm.AcceptCoin(coins.NewCoin(2.26, 17.91, 1.35))
			case "QUARTER":
				vm.AcceptCoin(coins.NewCoin(6.25, 24.26, 1.75))
			case "GET-COLA":
				vm.SelectProduct("COLA")
			case "GET-CHIPS":
				vm.SelectProduct("CHIPS")
			case "GET-CANDY":
				vm.SelectProduct("CANDY")
			case "COIN RETURN":
				vm.ReturnCoins()
			case "RESTOCK":
				vm.RestockProducts()
			case "RETRIVE-MONEY":
				vm.RetriveMoney()
			case "HELP":
				vm.ReturnCommands()
			default:
				fmt.Printf("command: '%s' not supported. Use HELP to see commands\n", c)
			}
		}

	}
}
