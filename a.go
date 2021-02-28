package main

import "fmt"

type Point struct {
	x, y int
}

func main() {

	q := &Point{10, 20}
	fmt.Println("Point before doStuff", *q)
	fmt.Println("Point before doStuff", q)

	doStuff(q)

	fmt.Println("Point after doStuff", *q)

}

func doStuff(p *Point) {
	p.x = 1000
}
