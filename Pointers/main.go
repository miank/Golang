package main

import "fmt"

func main() {
	var username string = "Go Developer"
	var usernamePointer *string = &username

	fmt.Println(*usernamePointer)
	fmt.Println(&usernamePointer)

	var age int = 12
	var isAwesome = true
	var height = 5.5

	// Declaring pointer of type

	var isAwesomePointer *bool = &isAwesome
	var heightPointer *float64 = &height
	var agePointer *int = &age

	fmt.Printf("agePointer: %v  ===> age value: %v \n", agePointer, *agePointer)
	fmt.Printf("isAwesomePointer: %v  ===> isAwesome value: %v \n", isAwesomePointer, *isAwesomePointer)
	fmt.Printf("heightPointer : %v  ===> height value: %v \n", heightPointer, *heightPointer)
}
