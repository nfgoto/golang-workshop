package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Привет, мир",
}

func Greet() {
	// Seed random number generator using the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number in the range of our list
	index := rand.Intn(len(helloList))

	// Call a function and receive multiple return values
	msg, err := greet(index)

	// Handle any errors
	if err != nil {
		log.Fatal(err)
	}

	// Print our message to the console
	fmt.Println(msg)
}

func greet(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// create an error, convert int type  to string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}

	return helloList[index], nil
}
