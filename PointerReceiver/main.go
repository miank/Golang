package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l *lion) breathe() {
	fmt.Println("Lion Breathes")
}

func (l *lion) walk() {
	fmt.Println("Lion Walks")
}

func main() {
	var a animal
	a = &lion{
		age: 5,
	}

	a.breathe()
	a.walk()

}
