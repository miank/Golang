package main

import "fmt"

func main() {
	slice1 := []int{}
	fmt.Println(slice1)
	fmt.Println(cap(slice1))

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)
}
