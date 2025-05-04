package main

import "fmt"

// mergeSortedSlices merges two sorted slices into one sorted slice
func mergeSortedSlices(slice1, slice2 []int) []int {
	merged := make([]int, 0, len(slice1)+len(slice2))
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			merged = append(merged, slice1[i])
			i++
		} else {
			merged = append(merged, slice2[j])
			j++
		}
	}

	merged = append(merged, slice1[i:]...)
	merged = append(merged, slice2[j:]...)

	return merged
}

func main() {
	slice1 := []int{1, 3, 5, 7}
	slice2 := []int{2, 4, 6, 8}

	result := mergeSortedSlices(slice1, slice2)
	fmt.Println(result)
}
