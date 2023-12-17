// Using struct instead of classes

package main

import (
	employee "hello/Struct/oop/employee"
)

func main() {
	employee.New("Sam", "Adolf", 30, 20)
	employee.LeavesRemaining()
}
