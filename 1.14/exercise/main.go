package main

import (
	"fmt"
	"time"
)

func main() {
	println("\nTo get to the value a pointer is associated to, you DEREFERENCE it")
	println("\nTo dereference a pointer, you use * in front of the variable name")
	println("\nNotice when declaring a pointer with the var statement, the * is in front of the type")
	println("\nDereferencing a nil pointer will cause an exception, you usually check if pointer is not nil before dereferencing it")
	println("\nYou don't always need to dereference a pointer to get to the value (for example with struct properties and functions)")
	println("\nTo get to the value a pointer is associated to, you use the * before the variable name")
	println("\nTo get to the value a pointer is associated to, you use the * before the variable name\n")

	var count1 *int
	count2 := new(int)
	countTemp := 4
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		// will cause a nil pointer exception, because no memory address allocated
		fmt.Printf("count 1: %#v \n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count 2: %#v \n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v \n", *count3)
	}
	if t != nil {
		// notice printed value does not start with a & so it is the value itself
		fmt.Printf("time (dereferenced): %#v \n", *t)
		fmt.Printf("time (String method): %#v \n", t.String())

	}

}
