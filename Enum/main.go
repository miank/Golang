package main

import "fmt"

const (
	smallSize = iota
	medium
	large
	extraSize
)

func main() {
	fmt.Println(smallSize)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraSize)
}
