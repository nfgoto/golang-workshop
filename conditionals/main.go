package main

import "fmt"

func main() {
	x, y := 3, 8

	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}

	if x > 2 && x < 8 {
		fmt.Println("x is between 2 and 8")
	}

	if x > 2 || x < 8 {
		fmt.Println("x is greater than 2 or less than 8")
	}

	a, b := 11.0, 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is greater than half of b")
	}

	///////////////////////
	// Switch
	///////////////////////

	z := 10

	switch z {
	case 10:
		fmt.Println("ten")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("many")
	}

	switch {
	case x > 1:
		fmt.Println("greater than 1")
	case x == 9:
		fmt.Println("equal to 9")
	case x != 7:
		fmt.Println("not 7")
	default:
		fmt.Println("whatever")
	}
}
