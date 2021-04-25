package main

import (
	"fmt"
)

func cleanUp(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func executeWorker() {
	// last executed even when panic or error in surrounding block
	defer cleanUp("A")
	// defer's are called in the reverse order of their appearance
	defer cleanUp("B")

	fmt.Println("Worker")
}

func main() {
	executeWorker()
}
