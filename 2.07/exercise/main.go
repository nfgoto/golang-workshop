package main

import (
	"fmt"
	"time"
)

func main() {
	// expressionless switch

	switch birthDay := time.Sunday; {
	case birthDay == time.Sunday || birthDay == time.Saturday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born on a weekday")
	}
}
