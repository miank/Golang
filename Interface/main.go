package main

import "fmt"

type animal interface {
	walk()
	breathe()
}
type mammal interface {
	feed()
}

type lion struct {
	age int
}

type dog struct {
	age int
}

// Non - struct Custom Type
type cat string

func (l lion) walk() {
	fmt.Println("Lion walk")
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) feed() {
	fmt.Println("Lion feeds young")
}

func (d *dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d *dog) walk() {
	fmt.Println("Dog walk")
}

func (c cat) breathe() {
	fmt.Println("Cat breathes")
}

func (c cat) walk() {
	fmt.Println("Cat walk")
}

func main() {
	var a animal
	var m mammal

	l := lion{
		24,
	}

	a = &l
	a.breathe()
	a.walk()
	m = l
	m.feed()

	a = &dog{age: 5}
	a.breathe()
	a.walk()

	a = cat("smokey")
	a.breathe()
	a.walk()
}
