package main

import "fmt"

func buildMessage() string {
	m := ""
	arr := [4]int{11, 22, 33, 44}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}

	return m
}

func main() {
	// creating array from condtant integer variable
	const lastIndex = 4
	arr := [lastIndex]string{}
	fmt.Println(arr)
	fmt.Println(buildMessage())
}
