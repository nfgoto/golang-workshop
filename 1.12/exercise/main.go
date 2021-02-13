package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		valid     bool
		count     int
		discount  float64
		name      string
		emails    []string
		startTime time.Time
	)

	fmt.Printf("Valid: %#v\n", valid)

	fmt.Printf("Name: %#v\n", name)

	fmt.Printf("Emails: %#v\n", emails)

	fmt.Printf("Discount: %#v\n", discount)

	fmt.Printf("Discount: %#v\n", count)

	fmt.Printf("Start time: %#v\n", startTime)

	// substitutions for formatted strings [template language]

	fmt.Printf("Simple value: %v \n", startTime)
	fmt.Printf("Value with extra data: %+v \n", startTime)
	fmt.Printf("Go syntax (equivalent to above): %#v \n", startTime)
	fmt.Printf("Type of value: %T \n", startTime)
	fmt.Printf("String value: %s \n", name)
	fmt.Printf("Decimal value: %d \n", count)

}
