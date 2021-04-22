package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println("iteration:", i)
	}

	fmt.Println("")

	for i := 0; i < 6; i++ {
		if i > 2 {
			break
		}
		fmt.Println("iteration:", i)
	}

	fmt.Println("")

	for i := 0; i < 6; i++ {
		if i < 2 {
			continue
		}
		fmt.Println("iteration:", i)
	}

	fmt.Println("")

	x := 5
	for x < 8 {
		fmt.Println(x)
		x++
	}

	fmt.Println("")

	y := 5
	for {
		if y > 10 {
			break
		}
		fmt.Println("not quite forever")
		y++
	}
}
