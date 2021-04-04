package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed]n", minArgs)
		os.Exit(1)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func findLonguestWord(args []string) string {
	var longuest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longuest) {
			longuest = args[i]
		}
	}
	return longuest
}

func main() {
	fmt.Printf("os.Args:\n%v\n\n", os.Args)
	if longuest := findLonguestWord(getPassedArgs(3)); len(longuest) > 0 {
		fmt.Println("Longuest command argument was:", longuest)
	} else {
		fmt.Println("Error occurred")
		os.Exit(1)
	}
}
