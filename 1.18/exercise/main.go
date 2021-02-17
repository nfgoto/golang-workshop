package main

import (
	"fmt"
)

var level = "package"

func main() {
	fmt.Println("Main start:", level)

	// function scope
	fmt.Println("\nIn Go, a child scope is defined by a pair of matching curly braces")
	fmt.Println("Parent-child relationship define at compilation")
	fmt.Println("When accessing a variable, Go looks at the scope it is defined in")
	println("Go looks for a variable in scope where used, if not found, looks in parent scope and so on  until package scope")
	println("Go uses static scope resolution - doesn't care where a function is called, but ONLY where it is defined, for scope resolution (contrary to /jS\n)")

	// create a shadow variable
	level := 42

	if true {
		fmt.Println("Block start:", level)
		funcA()
	}

	// block scope
	{
		level := "Nest 1"
		println("Nested start:", level)
	}

	println("Main end:", level)
}

func funcA() {
	println("funcA start:", level)
}
