package main

import "fmt"

func main() {
	var num int = 5

	// Prints the value stored in variable
	fmt.Println("Variable Value: ", num)

	// Prints the address of variable
	fmt.Println("Memory Address: ", &num)

	var ptr *int = &num
	fmt.Println("Pointer Value: ", *ptr)
	fmt.Println("Pointer Address: ", ptr)

	var name = "John"
	var ptr1 *string
	ptr1 = &name

	// Prints the value stored in variable
	fmt.Println("Variable Value: ", name)

	// Prints the address of variable
	fmt.Println("Memory Address: ", &ptr1)
	// Pointer Value
	fmt.Println("Pointer Value: ", *ptr1)

	var ptr3 *int
	ptr3 = &num
	num = 6
	*ptr3 = 7

	fmt.Println("Value of num ", num)
	fmt.Println("Value of ptr3 ", *ptr3)

	// Pointer and functions

	fmt.Println("The current value of num is ", num)

	update(&num)

	fmt.Println("After update pointer number is ", num)
}

func update(num *int) {
	*num = 30
}
