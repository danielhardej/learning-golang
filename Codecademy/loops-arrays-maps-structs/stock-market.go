package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

type Stock struct {
	name  string
	price float32
}

func (stock *Stock) updatePrice() {
	change := randomNumberGen(-10, 10)
	stock.price += change
}

func displayPrice(market []Stock) {
	for _, value := range market {
		fmt.Printf("%v: %.2f \n", value.name, value.price)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	stockMarket := []Stock{
		{"GOOG", 2313.50},
		{"AAPL", 157.28},
		{"FB", 203.77},
		{"TWTR", 50.00},
	}

	displayPrice(stockMarket)
	stockMarket[0].updatePrice()
	stockMarket[1].updatePrice()
	stockMarket[2].updatePrice()
	stockMarket[3].updatePrice()
	displayPrice(stockMarket)

}
