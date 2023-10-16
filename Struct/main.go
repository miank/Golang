package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

type point struct {
	x float64
	y float64
}

func main() {
	e := employee{}
	fmt.Println(e)
	emp := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}

	fmt.Printf("emp : %v\n", emp)
	fmt.Printf("emp : %+v\n", emp)
	fmt.Printf("emp : %#v\n", emp)
	fmt.Println(emp)

	n := emp.name
	fmt.Printf("Current Name is  : %+v\n", n)

	emp.name = "John"
	fmt.Printf("New Name is : %s\n", emp.name)

	// Creating struct without field names
	emp1 := employee{"John", 32, 4000}

	fmt.Printf("emp : %v\n", emp1)
	fmt.Printf("emp : %+v\n", emp1)
	fmt.Printf("emp : %#v\n", emp1)
	fmt.Println(emp)

	// Pointer to struct
	empP := &emp1
	fmt.Printf("EmpP: %v\n", empP)

	// Another way of creating pointer to struct
	empP1 := &employee{"John", 32, 4000}
	fmt.Printf("EmpP1: %v\n", empP1)

}
