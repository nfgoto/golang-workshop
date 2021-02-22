package main

import (
	"fmt"
	"time"
)

func main() {
	birthDay := time.Monday

	switch birthDay {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")
	}
}
