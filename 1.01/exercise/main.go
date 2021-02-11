package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// seed random number
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 5 (excluded)
	// add 1 to get a number between 1 and 5
	r := rand.Intn(5) + 1

	// repeat string * r times
	starts := strings.Repeat("*", r)

	fmt.Println(starts)
}
