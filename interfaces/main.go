package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func SumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// delaring an interface (collection of methods signatures)
type Shape interface {
	Area() float64
}

func main() {
	s1 := &Square{2}
	fmt.Println("s1:", s1)
	fmt.Println("s1 area:", s1.Area())

	c1 := &Circle{4}
	fmt.Println("c1:", c1)
	fmt.Println("c1 area:", c1.Area())

	// slice of structs that implement Shape interface
	shapes := []Shape{s1, c1}
	fmt.Println("Sum of areas of shapes:", SumAreas(shapes))
}
