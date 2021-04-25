package main

import (
	"fmt"
)

type Trade struct {
	Symbol string  // stock symbol
	Volume int     // number of shares
	Price  float64 // trade price
	Buy    bool    // true if buy trade, false if sell trade
}

func main() {
	// create object by field position in declaration
	t1 := Trade{"XYZ", 45, 11.76, true}
	fmt.Println("t1 is", t1)

	// print with field names
	fmt.Printf("t1:\t %+v\n", t1)

	// access field
	fmt.Printf("t1 price is %g\n", t1.Price)

	t2 := Trade{
		Symbol: "AAA",
		Volume: 100,
		Price:  5.43,
		Buy:    false,
	}
	fmt.Printf("t2:\t%+v\n", t2)

	// create object by omitting values, will be set to zero value
	t3 := Trade{
		Symbol: "BBB",
		Price:  234,
	}
	fmt.Printf("t3:\t%+v\n", t3)

}
