package main

import (
	"fmt"
	"math"
)

func brainwashString(saying *string) {
	*saying = "Beep Boop."
}

func brainwashBool(isIt *bool) {
	if *isIt {
		*isIt = false
	} else {
		*isIt = true
	}
}

func brainwashFloat(floater *float64) {
	*floater = 4.13
}

func main() {
	greeting := "Hello there!"
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)

	myBool := false
	brainwashBool(&myBool)
	fmt.Println("my bool is now:", myBool)

	myFloater := math.Pi
	brainwashFloat(&myFloater)
	fmt.Println("my floater is now:", myFloater)

}
