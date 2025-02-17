// Type assertions using maps in Golang
package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	// Print out the inventory
	for flavor, quantity := range donuts {
		fmt.Println(flavor, quantity)
	}

	// Add your code here
	firstChoiceCount := donuts["frosted"]
	fmt.Println(firstChoiceCount)

	secondChoiceCount, status := donuts["bavarian cream"]
	fmt.Println(secondChoiceCount)
	fmt.Println(status)
	if status {
		fmt.Println(secondChoiceCount)
	} else {
		fmt.Println("We do not sell that donut!")
	}

	// another (better) way to do it:
	secondChoice := "bavarian cream"
	if secondChoiceCount, ok := donuts[secondChoice]; ok {
		fmt.Println(secondChoiceCount)
	} else {
		fmt.Println("We do not sell that donut!")
	}
}
