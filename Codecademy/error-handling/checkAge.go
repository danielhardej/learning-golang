package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 18 {
		// Return a new error here
		return errors.New("Age must be at least 18")

	}
	return nil
}

func main() {
	age := 16

	// Check if an error exists
	err := checkAge(age)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid.")
	}

}
