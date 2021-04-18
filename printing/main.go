package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sentence := "this is a string literal"
	fmt.Println(sentence)

	sentence1 := `this is a:
	 raw string literal,
	 on multiple 
	 lines.`
	fmt.Println(sentence1)

	sentence2 := "strings are immutable in Go = trying to mutate them will cause a panic"
	println(sentence2)

	sentence3 := "this is %s\n"
	fmt.Printf(sentence3, "a formatted string")

	sentence4 := fmt.Sprintln("this is Sprintf argument")
	fmt.Printf("Output with details: %#v\n", sentence4)

	matrix := [][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("=================================================")
	for _, row := range matrix {
		for _, item := range row {
			fmt.Printf("|\t%d\t", item)
		}
		fmt.Println("|\n=================================================")
	}

	sentence5 := "this is another sentence"
	sentenceLength := len(sentence5)
	sentenceRuneCount := utf8.RuneCountInString(sentence5)
	fmt.Println("semtemce length:", sentenceLength)
	fmt.Println("semtemce rune count:", sentenceRuneCount)
	var reversed []byte
	for index := range sentence5 {
		revIndex := sentenceLength - 1 - index
		revChar := sentence5[revIndex]
		reversed = append(reversed, revChar)
	}
	fmt.Println("reversed: ", string(reversed))

	fmt.Println("Slice of bytes (unicode):", reversed)

	for _, char := range "karibu" {
		fmt.Printf("char %c as %#v\n", char, char)
	}

}
