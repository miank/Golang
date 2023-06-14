package main

import (
	"fmt"
	"math"
)

var globalVar = "test"

func main() {
	var aaa int
	fmt.Println(aaa)

	// Variable decleration with value
	var bbb int = 8
	fmt.Println(bbb)

	var a, b, c int
	fmt.Printf("a = %d, b = %d, c = %d \n ", a, b, c)

	var d, e int = 2, 4
	fmt.Println("The value of d = %d & e = %d \n", d, e)

	var (
		a1 int
		b1 float64
		c1 string = "a"
	)

	fmt.Printf("a1 = %d, b1 = %f, c = %s\n", a1, b1, c1)

	// Variable declared with no type

	var a3 = 20
	fmt.Printf("a3 value is %d ", a3)

	// Type interference
	t := 123
	fmt.Printf("Type is : %T \n", t)
	y := 3 + 5i
	fmt.Printf("Type is : %T \n", y)
	x := 'a'
	fmt.Printf("Type is : %T \n", x)
	w := true
	fmt.Printf("Type is : %T", w)

	result := math.Max(4, 5)
	fmt.Println(result)

	testGlobal()
}

func testGlobal() {
	fmt.Println("Global Variable Value is ", globalVar)
}
