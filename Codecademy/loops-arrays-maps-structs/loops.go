package main

import (
	"fmt"
)

func main() {

	lineBreak()
	fmt.Println("The classic for loop:")
	var count int
	for count = 1; count <= 12; count += 2 {
		fmt.Println(count)
	}

	lineBreak()
	fmt.Println("For loop as a while loop:")
	var guess int
	for guess != 56 {
		guess = ask()
		if guess == 56 {
			fmt.Println("You are correct! You may go through to the treasure!")
		}
	}

	lineBreak()
	fmt.Println("Showing off break and continue")
	for count := 0; count < 20; count++ {
		// WRITE CONTINUE STATEMENT TO SKIP IF COUNT EQUALS 8
		if count == 8 {
			fmt.Println("Skipping number", count)
			continue
		}
		// WRITE BREAK STATEMENT TO EXIT FOR LOOP IF COUNT EQUALS 15
		if count == 15 {
			fmt.Println("Exiting at number", count)
			break
		}
		fmt.Println(count)
	}

	lineBreak()
	fmt.Println("Looping through a map:")
	addressBook := map[string]string{
		"Jannet": "22 Water St",
		"Joe":    "241 North Rd",
		"Robert": "86 Stone St",
	}
	for name, address := range addressBook {
		fmt.Println(name, "lives at", address)
	}

	// Infinite loops
	// for {
	// 	count()
	// }
}

func ask() int {
	var input int
	fmt.Print("I am thinking of a number between 1-100: ")
	fmt.Scan(&input)
	return input
}

func count() {
	var input int
	var number int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)
	number += input
	fmt.Println("Total guests:", number)
}

func lineBreak() {
	fmt.Println("\n------------------------------------------------")
}
