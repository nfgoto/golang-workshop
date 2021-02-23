package main

import (
	"fmt"
)

func main() {
	wordCount := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	// struct type definition
	type mostFrequentWord struct {
		word  string
		count int
	}

	// create new struct instance
	mfw := mostFrequentWord{}

	for word, count := range wordCount {

		if mfw.count < count {
			mfw.word = word
			mfw.count = count
		}
	}

	fmt.Println("Most popular word:", mfw.word)
	fmt.Println("With of count of:", mfw.count)

}
