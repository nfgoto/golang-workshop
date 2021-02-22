package main

import (
	"fmt"

	"time"
)

func main() {
	birthDay := time.Monday

	// only the matching case is executed in GO
	switch birthDay {
	case time.Monday:
		fmt.Println("Born on", birthDay)
	case time.Tuesday:
		fmt.Println("Born on", birthDay)
	case time.Wednesday:
		fmt.Println("Born on", birthDay)
	case time.Thursday:
		fmt.Println("Born on", birthDay)
	case time.Friday:
		fmt.Println("Born on", birthDay)
	case time.Saturday:
		fmt.Println("Born on", birthDay)
	case time.Sunday:
		fmt.Println("Born on", birthDay)
	default:
		fmt.Println(birthDay, "is not aa valid day")

	}

}
