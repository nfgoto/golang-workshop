package main

import (
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Square struct {
	Center Point
	Length int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (s *Square) GetArea() int {
	return s.Length * s.Length
}

func NewSquare(x, y, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (received %d)", length)
	}

	square := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return square, nil
}

func main() {
	s1, err := NewSquare(3, 6, 2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("s1:%+v\n", s1)
	s1.Center.Move(-2, -1)
	fmt.Printf("s1:%+v\n", s1)
}
