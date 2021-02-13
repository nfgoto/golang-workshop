package main

import (
	"fmt"
)

func main() {
	visits := 15
	fmt.Println("Girst visit :", visits == 1)

	fmt.Println("Returning visit :", visits != 1)

	fmt.Println("Silver member :", visits < 21)

	fmt.Println("Platinum member :", visits > 30)
}
