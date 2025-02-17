// Error Handling in Go: Inventory Manager
// Welcome to the Inventory Manager project for SpeedyMart, designed to solidify our Go error handling skills. This exercise challenges us to complete a partially built inventory management application for a supermarket chain, focusing on implementing error-handling strategies.

// Our task is to complete the existing codebase by implementing the following features:

// - Adding items to the inventory
// - Removing items from the inventory
// - Updating item quantities
// - Helpful error handling

package main

import (
	"errors"
	"fmt"
)

type Item struct {
	name     string
	quantity int
}

type Inventory map[string]Item

const MaxInventorySize = 10

func addItem(inv Inventory, name string, quantity int) error {
	// TODO: Complete the addItem function
	if _, exists := inv[name]; exists {
		return fmt.Errorf("item already exists in inventory: %s", name)
	}
	if len(inv) > MaxInventorySize {
		// return fmt.Errorf("cannot add more items to inventory")
		return &InventoryFullError{inventorySize: len(inv)}
	} else {
		inv[name] = Item{name: name, quantity: quantity}
	}
	return nil
}

func wrapError(err error, message string) error {
	// TODO: Complete the wrapError function
	return fmt.Errorf("%s: %w", message, err)
}

func printErrorChain(err error) {
	// TODO: Complete the printErrorChain function
	if err == nil {
		return
	}
	fmt.Printf("Error: %v\n", err)
	printErrorChain(errors.Unwrap(err))
}

func removeItem(inv Inventory, name string) error {
	// TODO: Complete the removeItem function
	if _, itemExists := inv[name]; itemExists {
		delete(inv, name)
	} else {
		return wrapError(errors.New("Item does not exist"), fmt.Sprintf("invalid item '%s'", name))
	}
	return nil
}

type InventoryFullError struct {
	// TODO: Complete the InventoryFullError struct
	inventorySize int
}

func (e InventoryFullError) Error() string {
	// TODO: Complete the InventoryFullError Error method
	return fmt.Sprintf("%s could not be added: inventory full", e.inventorySize)
}

func updateItemQuantity(inv Inventory, name string, quantity int) {
	// TODO: Complete the updateItemQuantity function
	// fmt.Println("updating item quantity:", name, quantity)
	item, exists := inv[name]
	if !exists {
		panic(fmt.Sprintf("attempted to update non-existent item %s", name))
	}
	item.quantity = quantity
	inv[name] = item
}

func main() {
	inventory := make(Inventory)
	// TODO: Defer a function to recover from a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if err := addItem(inventory, "apple", 5); err != nil {
		fmt.Printf("Error adding apple: %v\n", err)
	} else {
		fmt.Println("Apple added to inventory")
		fmt.Println("Inventory:", inventory)
	}

	if err := addItem(inventory, "banana", 7); err != nil {
		fmt.Printf("Error adding banana: %v\n", err)
	} else {
		fmt.Println("Banana added to inventory")
		fmt.Println("Inventory:", inventory)
	}

	fmt.Println("\nRemoving non-existent item from inventory: orange")
	if err := removeItem(inventory, "orange"); err != nil {
		printErrorChain(err)
	}
	fmt.Println("Updating apple quantity to 10...")
	updateItemQuantity(inventory, "apple", 10)
	fmt.Println("Inventory:", inventory)

	fmt.Println("\nAdding 10 new empty items to inventory...")
	for i := range 9 {
		if err := addItem(inventory, string(i), 1); err != nil {
			fmt.Printf("Error adding new item: %v\n", err)
		}
	}
	fmt.Println("Inventory:", inventory)

	fmt.Println("\nAttempting to update non-existent item: orange")
	updateItemQuantity(inventory, "orange", 5)
}
