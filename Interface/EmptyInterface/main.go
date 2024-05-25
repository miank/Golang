package main

import "fmt"

type salaryCalc interface{}

func main() {
	var s salaryCalc
	fmt.Println(s)
}
