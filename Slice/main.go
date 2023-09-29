package main

import "fmt"

func main() {
	// declare a slice
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] // Creating a slice from an array
	fmt.Println(b)

	c := []int{6, 7, 8}
	fmt.Println(c)

	// changes done in slice is actually done in an array
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before ", darr)
	for i := range dslice {
		dslice[i]++
	}

	fmt.Println("array after", darr)

}
