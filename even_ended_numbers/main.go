package main

import "fmt"

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	// %q will display quoted string
	fmt.Printf("s is %q (type %T)\n", s, s)

	count := 0
	for a := 1000; a < 9999; a++ {
		for b := a; b < 9999; b++ {
			r := a * b
			c := fmt.Sprintf("%d", r)
			isEvenEnded := c[0] == c[len(c)-1]
			if isEvenEnded {
				count += 1
			}
		}
	}

	fmt.Printf("There are %d even-ended numbers by multiplying two 4-digit integers\n", count)

}
