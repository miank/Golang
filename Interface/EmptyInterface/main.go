package main

import "fmt"

func PrintAnything(v interface{}) {
	fmt.Println(v)
}

func main() {
	PrintAnything("Hello World !!")
}
