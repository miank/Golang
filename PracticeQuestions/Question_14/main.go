// Find the Most Frequent Element in an Array
// Problem Statement:

// Given an array of integers, return the element that appears the most number of times. If there are multiple such elements, return any one of them.

package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 1, 4, 1}

	m1 := make(map[int]int)

	max := 0
	num := -1
	for _, v := range arr {
		m1[v]++
		if m1[v] > max {
			max = m1[v]
			num = v
		}
	}

	fmt.Printf("The most frequent element in the array is %d and it appears %d times ", num, max)

}
