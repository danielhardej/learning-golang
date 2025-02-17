package main

import "fmt"

func main() {
	star := "Polaris"
	var starAddress *string = &star
	fmt.Println("The address of star is", starAddress)

	*starAddress = "Sirius"
	fmt.Println("The value of star is", star)

	// changing values with different scopes with pointers/references
	// take the global variable:
	var x int = 10
	// we won't be able to change the value of x in the main function
	addHundredByValue(x)
	fmt.Println("Value of x after addHundredByValue:", x)

	// but we can change the value of x in the main function by passing a pointer to x
	addHundredByReference(&x)
	fmt.Println("Value of x after addHundredByReference:", x)

	// another example:
	greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}

func addHundredByValue(num int) {
	num += 100
}

func addHundredByReference(numPtr *int) {
	*numPtr += 100
}

func brainwash(saying *string) {
	*saying = "Beep Boop"
}
