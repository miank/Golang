package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function")
	}()

	myFunc := func() {
		fmt.Println("This is another anonymous function")
	}

	// Named anonymous function
	add := func(a, b int) int {
		return a + b
	}

	result := add(3, 5)
	fmt.Println(result)

	myFunc()
}
