package main

import "fmt"

func mutate() (string, int, int) {
	return "updated", 100, 10
}

func notMutate(string, int, int) (string, int, int) {
	return "updated again", 10000, 55
}

func main() {
	a, b, c := 1, "two", true
	fmt.Println(a, b, c)

	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	query, limit, offset = "cat", 20, 1
	fmt.Println(query, limit, offset)

	query, limit, offset = mutate()
	fmt.Println(query, limit, offset)

	// primitive types passed by value
	// will not mutate original primitive-tyoed variables
	notMutate(query, limit, offset)
	fmt.Println(query, limit, offset)
}
