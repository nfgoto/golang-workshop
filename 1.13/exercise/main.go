package main

import (
	"fmt"
	"time"
)

func main() {
	println("Primitive values are passed by VALUE (copy) as function argumments\n")
	println("copies are placed on the STACK for simple memory management\n")
	println("A pointer is the directions to a value in memory, you follow them to get back to the value\n")
	println("No copy of the value is passed when using a pointer in a function\n")
	println("When using pointer, Go cannot manage memory using the stack because of scope no longer limited to the function\n")
	println("Go puts the value on the heap when using a pointer - allowing the value to exist until no parts of the software uses it anymore - these values are reclaimed during garbage collection (complex memory management)\n")
	println("Escape analysis = deciding when to put values on the heap\n")
	println("Memory management is not part of Go specification - it is internal implementation detail that can change at any time\n")
	println("Performance optimization approach: \n\t1 - do not prematurely optimize\n\t2 - when performance problem:\n\t\t1 - Mesure before change\n\t\t2 - Mesure after change\n")

	println("A pointer can be nil, when there is no value attached to it. Check for nil for using pointer.\n")

	println("You can compare pointers to each other. The comparison returns true ONLY when the pointer is equal to itself \n")

	// declare a pointer with var statement
	// the pointer will have a value of nil (absence of value) and NO addresss in memory
	var count1 *int

	// creating a pointer using built-in new function
	// gets memory for a type and returns a pointer to that address
	count2 := new(int)

	// create temporary variable to get pointer because not possible to get pointer of literal value
	countTemp := 5

	// create a pointer from an existing variable (with &)
	count3 := &countTemp

	// create a pointer to a complex value, here the Time struct
	// the value is a pointer to the Time struct with its properties' zero valued
	t := &time.Time{}

	// prints with %#v substitution: (*int)(nil)
	// reads pointer to a value of type int, no address in memory
	fmt.Printf("count1: %#v\n", count1)

	// prints with %#v substitution: (*int)(0xc0000140b8)
	// reads pointer to a value of type int with the address allocated in memory where to find the value
	fmt.Printf("count2: %#v\n", count2)

	// prints with %#v substitution: (*int)(0xc00010c008)
	// reads pointer to a value of type int with the address allocated in memory where to find the value
	fmt.Printf("count3: %#v\n", count3)

	// prints with %#v substitution: &time.Time{wall:0x0, ext:0, loc:(*time.Location)(nil)}
	// reads pointer to a value of type time.Time, notice starts with & so it is a pointer
	fmt.Printf("t: %#v\n", t)

}
