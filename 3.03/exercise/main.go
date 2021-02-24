package main

import (
	"fmt"
)

func main() {

	var a int8 = 125

	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++

		// will wrap around when the max value is exceeded
		// goes back to its lowest value, here -128 and 0

		fmt.Printf("%v ) int8 %v uint8 %v \n", i, a, b)

	}

}
