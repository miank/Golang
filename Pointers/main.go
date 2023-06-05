package main

import "fmt"

type Person struct {
	name string
	age  int
}

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

	// Return a pointer from function

	result := display()
	fmt.Println("The return value from function via pointer -->>", *result)

	// Call By Value
	callByValue(num)

	// Call By Reference
	callByReference(&num)

	person1 := Person{"John", 25}

	var ptr2 *Person // Create a pointer of type person
	ptr2 = &person1  // store the address of person

	fmt.Println("Person struct ", person1)

	fmt.Println("Person struct pointer ", ptr2)

	// access the name member
	fmt.Println("Name:", ptr2.name)

	// access the age member
	fmt.Println("Age:", ptr2.age)
}

func callByValue(num int) {
	num = 40
	fmt.Println("Call by Value ", num)
}

func callByReference(num *int) {
	*num = 50
	fmt.Println("Call by reference ", *num)
}

func update(num *int) {
	*num = 30
}

func display() *string {
	message := "Hello World"
	return &message
}
