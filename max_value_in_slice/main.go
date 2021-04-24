package main

import "fmt"

func main() {
	nums := []int{12, 34, 56, 78, 90, 32, 85, 93, 22}
	var max int = nums[0]

	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	fmt.Println("Max value in slice is/;", max)
}
