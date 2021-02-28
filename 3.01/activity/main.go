package main

import (
	"fmt"
)

type Item struct {
	name    string
	cost    float32
	taxRate float32
}

func main() {

	/*var message *string

	if message == nil {
		fmt.Println("error, unexpected nil value")
		return
	}

	fmt.Println(&message)
	*/

	items := []Item{
		Item{"cake", 0.99, 0.075},
		Item{"milk", 2.75, 0.015},
		Item{"butter", 0.87, 0.02},
	}

	var salesTaxTotal float32

	for _, item := range items {
		salesTaxTotal += item.cost * item.taxRate
	}

	fmt.Println("Sales Tax Total:", salesTaxTotal)
}
