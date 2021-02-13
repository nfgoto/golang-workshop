package main

import (
	"fmt"
)

func processRestaurantExpenses() {
	var total float64 = 2 * 13

	total = total + 4*1.43

	fmt.Println("Sub total is:", total, "euros")

	total = total + 4*1.43

	fmt.Println("Sub total is:", total, "euros")

	total = total - 5

	fmt.Println("Sub total is:", total, "euros")

	tip := total * 0.1

	fmt.Println("Tip is:", tip, "euros")

	total = total + tip

	fmt.Println("Total is:", total, "euros")

	split := total / 2

	fmt.Println("Split is:", split, "euros")
	println()
}

func processGrosseryReward() {
	visitCount := 24
	visitCount = visitCount + 1
	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you earned a reward")
	}
}

func processFullname() {
	givenName := "Sambo"
	firstName := "Juma"
	fullName := givenName + " " + firstName

	fmt.Println("Wilkommen", fullName)
}

func main() {
	processRestaurantExpenses()

	processGrosseryReward()

	processFullname()

}
