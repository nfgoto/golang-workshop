package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(num int) {
	num *= 2
}

// takes a pointer to an integer as argument
func doublePtr(num *int) {
	// dereference pointer and assign value
	*num *= 2
}

func main() {
	values := []int{4, 6, 11, 31}
	// complex types are passed by reference, changes inside func affect the passed argument
	doubleAt(values, 2)
	fmt.Println(values)

	num := 55
	double(num)
	// simple types are passed by value, a copy of the passed argument is used inside the gunc
	fmt.Println(num)

	num1 := 63

	// pointer represents an address in memory where the value is stored
	doublePtr(&num1) // create pointer from a variable => &myVar

	// the original variable is modified because the referenced value was modified in memory
	fmt.Println(&num1)
	fmt.Println(num1)
}
