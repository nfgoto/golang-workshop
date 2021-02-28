package main

import (
	"fmt"
)

func main() {
	// raw string
	post1 := `This is
a 
raw string`

	post2 := "This is an \n interpreted \t string"

	post3 := "Heute es ist windisch"

	println("\n=========== STRING ============\n\n")

	fmt.Print(post1, "\n\n")

	fmt.Print(post2, "\n\n")

	fmt.Print(post3, "\n\n")

	println("\n============ RUNE ===========\n\n")

	println(`A rune can store a single UTF-8 multi byte charqcter.

String literals are encoded using the UTF-8 encoding standard in Go.

The string type is not limited to UTG-8.

Different encodings use different numbers of bytes to encode text.

Legacy standards use 1 byte per character.

UTF-8 uses up to 4 bytes to encode a single character.

Go stores strings as a byte collection to allow for different encodings.

To perform operations on text of any encoding, FIRST convert it from a byte collection to rune collection.

To prevent bugs when you don't know the encoding of a string in advance:
- convert the string to UTF-8 
- UTF-8 is backward compatible with single-byte encoded text
`)

	username := "Señor_Bwana"

	fmt.Println("Accessing individual bytes of a strings:\n")

	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}

	println("\n\nConverting to rune collection:\n")

	runes := []rune(username)

	println(runes, "\n")

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}

	println()

	for offset, rune_ := range runes {
		fmt.Print("\nOffset: ", offset, " => ", string(rune_))
	}

	println("\n\nlooping safely over a string:\n")

	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println("Index:", index, ", as byte:", runeVal, ", as string:", string(runeVal))
	}

	println("\nchecklength of multi byte strings\n")

	name := "Übersetzen_Sie_bitte"

	println("Bytes:", len(name))

	println("Runes:", len([]rune(name)))

	println("\nlimit to 10 characters\n")

	println(string(name[:10]))

	println(string([]rune(name)[:10]))

}
