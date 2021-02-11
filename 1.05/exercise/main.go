package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// need to declare type to avoid incorrect type inference as int
	var seed int64 = 123456789
	rand.Seed(seed)
	fmt.Println(rand.Int())

	// short variable declarations
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
