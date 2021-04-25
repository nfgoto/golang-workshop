package main

import (
	"fmt"
	"strings"
)

func main() {
	wordMap := map[string]int{}

	text := `This is a line
	This is a line
	This is a line
	This is a line`

	words := strings.Fields(text)

	wordMap[words[0]] = 1

	for _, word := range words[1:] {
		// accessing non existing word returns 0 (zero value of this map values)
		wordMap[strings.ToLower(word)] += 1
	}

	fmt.Println("Word count:\n\t", wordMap)
}
