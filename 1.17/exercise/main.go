package main

import (
	"fmt"
)

// creating an enum
const (
	first = iota

	second

	third

	fourth
)

func main() {
	fmt.Println("First:", first)
	fmt.Println("Second:", second)
	fmt.Println("Third:", third)
	fmt.Println("Fourth:", fourth)
}
