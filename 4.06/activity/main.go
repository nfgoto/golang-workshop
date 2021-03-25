package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		// modifying array element in place
		arr[i] = i + 1
	}
	return arr
}

func square(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] *= arr[i]
	}
	return arr
}

func main() {
	var arr [10]int
	arr = fillArray(arr)
	arr = square(arr)
	fmt.Println(arr)
}
