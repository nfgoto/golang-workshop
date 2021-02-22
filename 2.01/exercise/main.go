package main

import (
	"fmt"
)

func main() {
	input := 5

	// can only be used in function scope
	if input%2 == 0 {

		fmt.Println(input, "is even")
	}

	if input%2 == 1 {

		fmt.Println(input, "is odd")
	}

}
