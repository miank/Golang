// Write a function in Go that:

// Takes a slice of integers as input.

// Removes all duplicate elements.

// Returns a new slice containing only unique elements, preserving their first appearance order.

package main

import "fmt"

func FindDuplicates(nums []int) []int {
	// Your code here

	m1 := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		m1[nums[i]]++
	}

	result := []int{}

	for num, count := range m1 {
		if count > 1 {
			result = append(result, num)
		}
	}

	return result

}

func main() {
	ss := []int{1, 2, 3, 2, 4, 5, 1}
	fmt.Println(FindDuplicates(ss))

}
