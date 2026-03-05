package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// rectangle
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Height * r.Width }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Height + r.Width) }

// Circle
type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Write TotalArea(shapes []Shape) float64 — sum the areas of all shapes
// This function should work for any combination of shapes without knowing their concrete types.
func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()

	}
	return float64(total)
}

func main() {
	results := []Shape{
		Rectangle{Width: 4, Height: 65},
		Circle{Radius: 33.0},
	}

	for _, s := range results {
		fmt.Println(s)
	}

	fmt.Println("Total area: ", TotalArea(results))
}
