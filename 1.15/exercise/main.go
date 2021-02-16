package main

import (
	"fmt"
)

func addFive(count int) {
	count += 5
	println("\nAdded five", count)
}

// function accepting a pointer as argument
func addFivePoint(num *int) {
	// derefence the pointer to work with the valie itself
	*num += 5
	println("\nAdded five with pointer", *num)
}

func main() {
	var count int

	addFive(count)
	fmt.Println("\nCount post addFive", count)

	// you need the & to get the count pointer
	addFivePoint(&count)
	fmt.Println("\nCount post addFivePoint", count)
}
