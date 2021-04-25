package main

import (
	"fmt"
	"net/http"
)

func getResponseContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // make sure to close response body

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("cannot find Content-Type header")
	}

	return ctype, nil
}

func main() {
	url := "https://golang.org"

	ctype, err := getResponseContentType(url)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println("Response Content-Type header is:", ctype)
	}

}
