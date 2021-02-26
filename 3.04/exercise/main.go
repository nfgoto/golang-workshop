package main

import (
	"fmt"

	"math"

	"math/big" // dealing with numbers bigger/lower thab core types
)

func main() {
	intA := math.MaxInt64

	// wrap around, will have the lowest value for int84
	intA = intA + 1

	// create big int

	bigA := big.NewInt(math.MaxInt64)

	// adding 1 to big int

	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)

	fmt.Println("Int   :", intA)

	fmt.Println("Big Int : ", bigA.String())

	// byte is an alias for uint8, used when dealing with I/O
	var bb byte = 24

	fmt.Printf("%v == %#v \n", bb, bb)

}
