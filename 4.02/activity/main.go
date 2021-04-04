package main

import (
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	m["305"] = "Sue"
	m["204"] = "Bob"
	m["631"] = "Jake"
	m["073"] = "Tracy"

	arg := os.Args[2]
	if arg == "" {
		fmt.Println("You must enter an argument")
		os.Exit(1)
	}
	if m[arg] == "" {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	fmt.Println("Hi,", m[arg])

}
