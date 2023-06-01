package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("The lion breathes")
}

func (l lion) walk() {
	fmt.Println("The lion walks ")

}

func main() {
	var a animal
	a = lion{age: 12}
	a.breathe()
	a.walk()
}
