package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string  // stock symbol
	Volume int     // number of shares
	Price  float64 // trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// trade factory / constructor with input validation
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0 {
		return nil, fmt.Errorf("price must be >= 0 (was %g)", price)
	}

	// create pointer to Trade struct
	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
}

func main() {
	t1, err := NewTrade("CCC", 23, 65.1, true)

	if err != nil {
		fmt.Println("Error at construction:", err)
		os.Exit(1)
	}

	fmt.Printf("t1 is %+v\n", t1)

}
