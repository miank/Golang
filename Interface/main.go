package main

import "fmt"

type animal interface {
	walk()
	breathe()
}

type lion struct {
	age int
}

type dog struct {
	age int
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l dog) breathe() {
	fmt.Println("Dog breathes")
}

func (l dog) walk() {
	fmt.Println("Dog walk")
}

func main() {
	ll := lion{
		24,
	}

	var a animal
	a = ll
	a.breathe()
	a.walk()

	a = dog{age: 5}
	a.breathe()
	a.walk()
}
