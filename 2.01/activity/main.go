package main

import (
	"fmt"
)

const lowerBoundary = 1

const upperBoundary = 100

func main() {
	for input := lowerBoundary; input <= upperBoundary; input++ {
		if input%15 == 0 {
			fmt.Println(input, ": FFizzBuzz")
		} else if input%3 == 0 {
			fmt.Println(input, ": Fizz")
		} else if input%5 == 0 {
			fmt.Println(input, ": Buzz")
		}
	}
}
