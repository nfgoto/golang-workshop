package main

import (
	"fmt"
)

// declaring a named struct type
type Trade struct {
	Symbol string  // stock symbol
	Volume int     // number of shares
	Price  float64 // trade price
	Buy    bool    // true if buy trade, false if sell trade
}

// return the trade value
func (t *Trade) getValue() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

type Point struct {
	// compact field declaration (same type - prefer full declaration)
	X, Y int
}

// need to use a pointer to the struct
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	// create a struct (without specifying field names)
	t := Trade{"AAA", 100, 5.43, false}

	fmt.Printf("t:\t%v\n", t)
	fmt.Println("trade value is:", t.getValue())

	// create a struct (specifying field names)
	t1 := Trade{
		Symbol: "BBB",
		Volume: 23,
		Price:  9.12,
		Buy:    true,
	}

	fmt.Printf("t1:\t%+v\n", t1)
	fmt.Println("trade value is:", t1.getValue())

	p1 := Point{1, 2}
	p1.Move(3, 4)
	fmt.Printf("p1 is %+v\n", p1)
}
