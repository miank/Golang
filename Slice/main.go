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
	numa := [3]int{78, 79, 80}

	nums1 := numa[:] // contains all the elements of the slice
	nums2 := numa[:]
	fmt.Println("Array before change 1 ", nums1)
	fmt.Println("Array before change 1 ", nums2)
	nums1[0] = 100
	fmt.Println("Array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("Array after modification to slice nums2", numa)

	// Length and Capacity of slice
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice is %d and capacity is %d \n", len(fruitSlice), cap(fruitSlice))
	fruitSlice = fruitSlice[:cap(fruitSlice)]
	fmt.Println("After re-slicing length is ", len(fruitSlice),
		"and capacity is ", cap(fruitSlice))

	// Creating a slice using make
	l := make([]int, 5, 5)
	fmt.Println(l)

}
