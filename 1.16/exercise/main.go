package main

import (
	"fmt"
)

const (
	GlobalLimit = 500

	MaxCacheSize = 10 * GlobalLimit

	CacheKeyBook = "book_"

	CacheKeyCd = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, value string) {
	if len(cache) >= MaxCacheSize {
		return
	}
	cache[key] = value
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func SetBook(isbn, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCd + sku)
}

func SetCD(sku, name string) {
	cacheSet(CacheKeyCd+sku, name)
}

func main() {
	// create a map
	cache = make(map[string]string)

	SetBook("1234-5678", "The Go Workshop")

	SetCD("2345-6789", "The Go Workshop Audiobook")

	fmt.Println("Book:", GetBook("1234-5678"))

	fmt.Println("CD:", GetCD("2345-6789"))

}
