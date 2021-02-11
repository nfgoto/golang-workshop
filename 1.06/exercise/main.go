package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// multiple short variable declarations
	Debug, LogLevel, startUpTime := getConfig()

	fmt.Println(Debug, LogLevel, startUpTime)

	// multiple long variable declaration of same type
	var start, middle, end float32
	fmt.Println(start, middle, end)

	// multiple long variable declaration and init of different types
	var name, left, right, top, bottom = "Toto", 1, 1.5, 2, 2.5
	fmt.Println(name, left, right, top, bottom)

	var Debug1, LogLevel1, startUpTime1 = getConfig()
	fmt.Println(Debug1, LogLevel1, startUpTime1)

}
