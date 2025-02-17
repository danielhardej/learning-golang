package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

// WRITE CONTAINS FUNCTION BELOW
func contains(menu []string, order string) bool {
	for _, item := range menu {
		if item == order {
			return true
		}
	}
	return false
}

func showMenu(menu []string) {
	for _, item := range menu {
		fmt.Println("-", item)
	}
}

func main() {
	fastfoodMenu := []string{"Burgers", "Nuggets", "Fries", "Chicken Tendies", "Steak", "Ribs"}
	fmt.Println("The fast food menu has these items: ")
	showMenu(fastfoodMenu)

	var total int
	var order string
	// WRITE INDEFINITE LOOP ASKING FOR ORDERS BELOW
	for order != "quit" {
		order = askOrder()
		if order != "quit" {
			if contains(fastfoodMenu, order) {
				total += 4
				fmt.Println(order, "added. Would you like anything else?")
			} else {
				fmt.Println("Sorry, that item is not on the menu. Please choose from: ")
				showMenu(fastfoodMenu)
			}
		}
	}

	var twoDollarBillCount int
	// WRITE DEFINITE LOOP COUNTING $2 BILLS BELOW
	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("The amount left to pay is:", amount)
		fmt.Println("Handing over a $2 bill...")
		twoDollarBillCount += 1
	}

	fmt.Println("The total for the order is", total)
	fmt.Println("We paid with", twoDollarBillCount, "$2 bills!")
}
