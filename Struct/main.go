package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println("Structs in Golang")
	user := User{
		Name:   "Hitesh",
		Email:  "Hitesh@gmail.com",
		Status: true,
		Age:    22,
	}
	fmt.Printf("User Details are %+v: \n", user)
	fmt.Printf("User Name is %v, Email is %v and age is %v: \n", user.Name, user.Email, user.Age)

	fmt.Println(&person{
		name: "Fred",
	})

	fmt.Println(newPerson("John"))

	s := person{
		name: "Sean",
		age:  50,
	}
	fmt.Println(s)

	sp := s
	sp.age = 22
	fmt.Println(s.age)
	
}
