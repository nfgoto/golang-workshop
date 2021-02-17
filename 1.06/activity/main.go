package main

func main() {
	count := 0

	if count < 5 {
		count = 10

		count++
	}

	println(count == 11)
}
