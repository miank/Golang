package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function")
	}()

	myFunc := func() {
		fmt.Println("This is another anonymous function")
	}

	myFunc()
}
