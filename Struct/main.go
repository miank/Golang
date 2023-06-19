package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
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
	fmt.Printf("User Name is %v, Email is %v and age is %v: ", user.Name, user.Email, user.Age)

}
