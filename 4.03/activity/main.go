package main

import (
	"fmt"
	"os"
)

type Locale struct {
	language string
	country  string
}

func main() {
	supportedLocales := [5]Locale{
		{language: "en", country: "US"},
		{language: "en", country: "CN"},
		{language: "fr", country: "CN"},
		{language: "fr", country: "FR"},
		{language: "ru", country: "RU"},
	}
	inputLanguage := os.Args[1]
	inputCountry := os.Args[2]
	if inputLanguage == "" || inputCountry == "" {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	inputLocale := Locale{language: inputLanguage, country: inputCountry}
	isSupported := false
	for i := 0; i < len(supportedLocales); i++ {
		current := supportedLocales[i]
		if current.language == inputLocale.language && current.country == inputLocale.country {
			isSupported = true
		}
	}
	if !isSupported {
		fmt.Printf("Locale not supported: %s_%s\n", inputLocale.language, inputLocale.country)
		os.Exit(1)
	}

}
