// Problem Statement:

// You are building a small Go application that calculates the area and perimeter of basic shapes.
// Split the logic into two packages:

// shapes package:

// Define two exported functions:

// AreaOfRectangle(length, width float64) float64

// PerimeterOfRectangle(length, width float64) float64

// main package:

// Import the shapes package.

// Call the functions with sample values (e.g., length = 10, width = 5) and print the results.

package main

import (
	"fmt"
	"hello/PracticeQuestions/PackageModularizationAndExport/shapes"
)

func main() {

	length := 10.0
	width := 5.0

	area := shapes.AreaOfRectangle(length, width)
	perimeter := shapes.PerimeterOfRectangle(length, width)

	fmt.Printf("Area: %.2f\n", area)
	fmt.Printf("Perimeter: %.2f\n", perimeter)

}
