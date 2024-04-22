package main

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	ints := []int{1, 2, 3}
	strings := []string{"One", "Two", "Three"}
	Print(ints)
	Print(strings)
}
