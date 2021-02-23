package main

import (
	"fmt"
)

func main() {
	numbers := []int{4, 2, 9, 6, 1, 3, 5}

	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				temp := numbers[j]

				numbers[j] = numbers[i]

				numbers[i] = temp
			}

		}
	}

	fmt.Println(numbers)

}
