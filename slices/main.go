package main

import "fmt"

func main() {
	langs := []string{"go", "js", "rust", "bash", "python", "c++", "typescript"}

	fmt.Printf("langs are %v (type %T)\n", langs, langs)

	fmt.Printf("length is %v (type %T)\n", len(langs), len(langs))

	fmt.Printf("langs[2] is %v (type %T)\n", langs[2], langs[2])

	fmt.Printf("langs[:2] is %v (type %T)\n", langs[:2], langs[:2])

	fmt.Printf("langs[2:] is %v (type %T)\n", langs[2:], langs[2:])

	fmt.Printf("langs[4:6] is %v (type %T)\n", langs[4:6], langs[4:6])

	fmt.Printf("length of langs[4:6] is %v (type %T)\n", len(langs[4:6]), len(langs[4:6]))

	for i := 0; i < len(langs); i++ {
		fmt.Println("\t", langs[i])
	}
	fmt.Println("")

	for index := range langs {
		fmt.Printf("(%d)\t", index)
	}
	fmt.Println("")

	for _, value := range langs {
		fmt.Printf("%s\t", value)
	}
	fmt.Println("")

	for _, value := range append(langs, "lua", "crystal", "kotlin") {
		fmt.Printf("%s\t", value)
	}
	fmt.Println("")
}
