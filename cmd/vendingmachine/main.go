package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TarasLykhenko/vending-machine-kata/internal/coins"
	"github.com/TarasLykhenko/vending-machine-kata/internal/vendingmachine"
)

func main() {
	vm := vendingmachine.NewVendingMachine()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(vm.Display())
		fmt.Print("Input: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		commands := strings.Split(text, ",")

		for _, c := range commands {
			switch c {
			case "PENNY":
				coins.NewCoin(2.5, 19.05, 1.52)
			case "NICLE":
				coins.NewCoin(5, 21.21, 1.95)
			case "DIME":
				coins.NewCoin(2.26, 17.91, 1.35)
			case "QUORTER":
				coins.NewCoin(6.25, 24.26, 1.75)
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
			default:
				fmt.Printf("command: '%s' not supported.\n", c)
			}
		}

	}
}
