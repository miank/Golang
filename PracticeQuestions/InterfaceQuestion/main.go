// Question: Write a Go program that defines an interface Shape with a method Area() float64.
// Implement this interface for two types: Circle (with a radius) and Rectangle (with width and height).
// Additionally, create a function that calculates the total area of a slice of shapes, handling cases
// where a shape might be invalid (e.g., negative dimensions) by returning an error.

package main

import (
	"errors"
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	if c.Radius < 0 {
		return 0.0
	}
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	if r.Width < 0 || r.Height < 0 {
		return 0.0
	}
	return r.Width * r.Height
}

func TotalArea(shapes []Shape) (float64, error) {
	var total float64
	for _, shape := range shapes {
		area := shape.Area()
		// Check for invalid shapes by inspecting their type and dimensions
		switch s := shape.(type) {
		case Circle:
			if s.Radius < 0 {
				return 0.0, errors.New("invalid shape: negative radius in Circle")
			}
		case Rectangle:
			if s.Width < 0 || s.Height < 0 {
				return 0.0, errors.New("invalid shape: negative dimensions in Rectangle")
			}
		}
		total += area
	}
	return total, nil
}

func main() {
	shapes := []Shape{
		Circle{Radius: 2.0},
		Rectangle{Width: 3.0, Height: 4.0},
		Circle{Radius: -1.0},
		Rectangle{Width: 5.0, Height: 2.0},
	}

	// Calculate total area and handle errors
	total, err := TotalArea(shapes)
	if err != nil {
		fmt.Printf("Total Area: %.2f, Error: %v\n", total, err)
	} else {
		fmt.Printf("Total Area: %.2f\n", total)
	}
}
