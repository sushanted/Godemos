package main

import "fmt"

type Shape interface {
	Area() float64
}

type Line struct{}

func (line Line) Area() float64 {
	return 0.0
}

type Circle float64

func (radius Circle) Area() float64 {
	return float64(radius * radius * (22.0 / 7.0))
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func main() {
	shapes := []Shape{Line{}, Circle(7), Triangle{3, 6}}

	for _, shape := range shapes {
		fmt.Printf("Area of %T is %v\n", shape, shape.Area())
	}
}
