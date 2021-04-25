package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"NFLX": 505.55,
		"GOOG": 2315.305,
		"INTC": 59.24,
	}

	fmt.Printf("stocks are %v (type %T)\n", stocks, stocks)
	fmt.Printf("length is %d\n", len(stocks))
	fmt.Printf("stocks[\"INTC\"] is %v\n", stocks["INTC"])
	fmt.Printf("zero value is %v\n", stocks["TOTO"])

	// check if in map
	value, ok := stocks["ALI"]
	if !ok {
		fmt.Println("ALI not found")
	} else {
		fmt.Println(value)
	}

	// set
	stocks["TOTO"] = 234.87
	fmt.Println(stocks)

	// delete
	delete(stocks, "GOOG")
	fmt.Println(stocks)

	// loop
	for key, value := range stocks {
		fmt.Printf("(%s,%g)\n", key, value)
	}
}
