package main

import "fmt"

func main() {
	book := "this is a sequence of bytes"
	fmt.Println(book)

	fmt.Printf("Its length: %v bytes\n", len(book))

	// uint8 is a byte
	fmt.Printf("book[0] is %v (type %T)\n", book[0], book[0])

	// strings are immutable in Go
	// book[2] = 'b'

	// slice a string, [start:end], end exclusive (half-empty)
	fmt.Println(book[5:10])

	// slice a string, [start:]
	fmt.Println(book[5:])

	// slice a string, [:end]
	fmt.Println(book[:5])

	// consatenate strings
	fmt.Println("hello" + book)

	// unicode
	fmt.Println("I accept unicode: ㋛")

	fmt.Println(`this
	is
	a
	multi-line
	raw
	string`)

}
