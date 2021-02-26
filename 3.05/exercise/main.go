package main

import (
	"fmt"
)

func main() {
	// raw string
	post1 := `This is
a 
raw string`

	post2 := "This is an \n interpreted \t string"

	post3 := "Heute es ist windisch"

	fmt.Print(post1, "\n\n")

	fmt.Print(post2, "\n\n")

	fmt.Print(post3, "\n\n")

}
