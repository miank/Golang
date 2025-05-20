package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 1, 4, 1}

	m1 := make(map[int]int)

	for _, v := range arr {
		m1[v]++
	}

	max := 0
	num := 0
	for k, v := range m1 {
		if v > max {
			max = v
			num = k
		}
	}

	fmt.Printf("The most frequent element in the array is %d and it appears %d times ", num, max)

}
