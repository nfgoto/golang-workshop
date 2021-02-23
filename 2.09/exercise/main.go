package main

import (
	"fmt"
)

func main() {
	// looping over slice of strings

	names := []string{"toto", "titi", "tutu"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
