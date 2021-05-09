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

	// anonymous struct
	anon := struct {
		name   string
		field1 int
		field2 string
	}{"anon", 987, "field 2"}

	fmt.Printf("anon:\t%+v\n", anon)

	// pointer to a struct

	type Executive struct {
		firstName string
		lastName  string
		position  int
		age       int
	}

	e1Ptr := &Executive{firstName: "Toto", lastName: "Goto", position: 1, age: 37}

	fmt.Println("e1 last name:", (*e1Ptr).lastName)
	// no need to dereference pointer to access fiellds
	fmt.Println("e1 first name:", e1Ptr.firstName)

	type A struct {
		// anonymous fields (name is same as type internally)
		string
		int
	}
	a1 := A{"Ok", 0}
	fmt.Println("a1:", a1)
	fmt.Println("a1 string:", a1.string)

	type B struct {
		a  A
		f1 bool
		f2 rune
	}

	b1 := B{
		// nested struct
		a:  A{"KO", 1},
		f1: false,
		f2: 'n',
	}
	fmt.Printf("b1: %+v\n", b1)
	fmt.Printf("b1.a.int: %+v\n", b1.a.int)

	// nested struct as anonymous field
	type C struct {
		f3 float64
		B
	}
	c1 := C{
		f3: 3.45677,
		B: B{
			a:  A{"yes", 1},
			f1: false,
			f2: 'y',
		},
	}
	fmt.Printf("c1: %+v\n", c1)
	fmt.Printf("f1: %+v\n", c1.f1)
}
