package main

import "fmt"

func main() {

	f1()
}

func f1() {
	defer fmt.Println("Defer in f1")
	f2()
	fmt.Println("After panic in f1")
}

func f2() {
	defer fmt.Println("Defer in f2")
	panic("Panic Demo")
	fmt.Println("After panic in f2")
}
