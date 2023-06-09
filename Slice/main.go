package main

import "fmt"

func main() {
	// declare a slice
	s := []int{1, 2}
	fmt.Println("Slice :", s)
	fmt.Printf("Slice length %d and Capacity %d: ", len(s), cap(s))

	letters := []string{"a", "b", "c"}
	fmt.Println("Slice is ", letters)
	fmt.Println("Length of slice is ", len(letters))
	fmt.Println("Capacity of slice is ", cap(letters))

	numbers := [5]int{1, 2, 3, 4, 5}
	// Creating a slice from existing slice
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length  = %d \n", len(num1))
	fmt.Printf("Capacity length = %d \n", cap(num1))

	num2 := numbers[2:]
	fmt.Println("Only start")
	fmt.Printf("num2=%v\n", num2)
	fmt.Printf("length  = %d \n", len(num2))
	fmt.Printf("Capacity length = %d \n", cap(num2))

	num3 := numbers[:3]
	fmt.Println("Only end")
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("length  = %d \n", len(num3))
	fmt.Printf("Capacity length = %d \n", cap(num3))

}
