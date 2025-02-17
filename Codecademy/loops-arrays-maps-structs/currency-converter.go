package main

import (
	"fmt"
)

func main() {

	var currencies = map[string]float32{
		// 1 USD equals...
		"JPY": 130.2,
		"EUR": 0.95,
		"GBP": 0.82,
		"AUD": 1.5,
	}

	var dollarAmount float32
	var currency string
	var convertedAmount float32

	fmt.Println("Enter how many dollars you want to convert: ")
	fmt.Scan(&dollarAmount)

	if dollarAmount <= 0 {
		fmt.Println("Invalid amount.")
	}

	fmt.Println("Enter the currency you want to convert to convert:")
	for currency, _ := range currencies {
		fmt.Println("-", currency)
	}
	fmt.Scan(&currency)

	if exchangeRate, isValid := currencies[currency]; isValid {
		convertedAmount = dollarAmount * exchangeRate
	} else {
		fmt.Println("Invalid currency:", currency)
		convertedAmount = 0
	}

	fmt.Println("Converted amount:", convertedAmount)

}
