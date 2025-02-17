// An example of panic and recover in Go to handle exceptions
// without interrupting the program.

package main

import "fmt"

func accessSlice(slice []int, index int) {
	// Recover out-of-bounds access here
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic:", r)
		}
	}()
	fmt.Println("Attempting to access slice...")
	fmt.Println(slice[index])
}

func main() {
	nums := []int{1, 2, 3}

	accessSlice(nums, 5)

	fmt.Println("Accessing a valid index:", nums[1])
}
