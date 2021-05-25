package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}
type square struct {
	side1, side2 float64
}

func main() {
	t := triangle{
		base:   10,
		height: 15,
	}
	s := square{
		side1: 20,
		side2: 18,
	}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side1 * s.side2
}
