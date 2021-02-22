package main

import (
	"fmt"
)

func main() {
	input := 2

	if input%2 == 1 {
		fmt.Println(input, "is odd")
	} else {
		fmt.Println(input, "is even")
	}

}
