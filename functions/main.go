package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	x, y := 99, 45
	fmt.Printf("%d + %d = %d\n", x, y, add(x, y))

	whole, remainder := divide(x, y)
	fmt.Printf("%d / %d = %dr%d\n", x, y, whole, remainder)
}
