package main

import (
	"fmt"
)

func defineArray() [10]int {
	var arr [10]int
	return arr
}

func main() {
	arr := defineArray()
	fmt.Printf("%#v\n", arr)

	// dynamically generated array
	// length set at compile time based on initialization
	// fixed length at runtime

	arr1 := [...]int{0, 1, 2, 3, 4}
	fmt.Printf("%#v\n", arr1)

}
