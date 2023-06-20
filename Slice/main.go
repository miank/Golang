package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// declare a slice
	s := []int{1, 2}
	fmt.Println("Slice :", s)
	fmt.Printf("Slice length %d and Capacity %d \n : ", len(s), cap(s))

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

	// Refering to the original slice only

	numbers[3] = 8

	num3 := numbers[:3]
	fmt.Println("Only end")
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("length  = %d \n", len(num3))
	fmt.Printf("Capacity length = %d \n", cap(num3))

	num4 := numbers[:]
	fmt.Println("Only end")
	fmt.Printf("num3=%v\n", num4)
	fmt.Printf("length  = %d \n", len(num4))
	fmt.Printf("Capacity length = %d \n", cap(num4))

	// Creating a slice  using a make function
	numbers1 := make([]int, 3, 5)
	fmt.Printf("numbers1 = %v\n", numbers1)
	fmt.Printf("length of numbers1 = %d\n", len(numbers1))
	fmt.Printf("capacity of numbers1 = %d\n", cap(numbers1))

	// Slice iteration
	letters1 := []string{"a", "b", "c", "d", "e", "f"}

	for i, letter := range letters1 {
		fmt.Printf("%d %s \n", i, letter)
	}

	for _, letter := range letters1 {
		fmt.Printf("%s \n", letter)
	}

	// Without value
	i := 0
	for range letters {
		fmt.Println(i)
		i++
	}

	// Find an item in slice

	source := []string{"san", "man", "tan"}
	result := findInSlice(source, "san")
	fmt.Println(result)

	// Convert slice to json

	s1 := []string{"Hello", "World"}

	j, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Printf(string(j) + "\n")
	}

	// FindAndDelete -- need to do

	// Sorting of numbers in slice

	nums := []int{1, 2, 3, 4}

	sortSlice(nums) // to do

	// Copy an slice
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5) // Another way to make a slice
	numberOfElementsCopied := copy(dst, src)
	fmt.Printf("Number of Elements Copied: %d \n", numberOfElementsCopied)
	fmt.Printf("src %v \n", src)
	fmt.Printf("dst %v \n", dst)

	// Find Item in a Slice
	source1 := []string{"san", "man", "tan"}
	result1 := findInSlice(source1, "san")
	fmt.Printf("Item was found : %t\n", result1)

	// Create a slice of string
	var ss []string
	ss = make([]string, 4)
	// appending the slice
	ss = append(ss, "abc")
	ss = append(ss, "def", "klm")
	ss = append(ss, "ghi")

	fmt.Println("Looping through the slice")
	for _, c := range ss {
		fmt.Println(c)
	}

}

func findInSlice(s []string, value string) bool {
	for _, item := range s {
		if item == value {
			fmt.Printf("Value %s was found at index %s \n", value, item)
			return true
		}
	}
	return false
}

func sortSlice(s []int) {

}
