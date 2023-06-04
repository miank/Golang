package main

import "fmt"

func mult(a, b int) int {
	result := a * b
	fmt.Println("Result :", result)
	return result
}
func main() {
	fmt.Println("Start")

	// Multiple defer statements executes in LIFO Order
	defer fmt.Println("End")
	defer mult(34, 56)
	defer mult(10, 10)
}
