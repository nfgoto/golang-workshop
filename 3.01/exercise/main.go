package main

import (
	"fmt"
	"unicode"
)

func passwordCker(pw string) bool {
	// convert pw into slice of Rune, safe for utf8
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}

	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {

	uPw := "Not-saFe"

	sPw := "5*o/Ti#g56i9l8"

	fmt.Printf("\"%s\" safe password : %v\n", uPw, passwordCker(uPw))

	fmt.Printf("\"%s\" safe password : %v\n", sPw, passwordCker(sPw))

}
