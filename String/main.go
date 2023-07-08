package main

import "fmt"

func main() {
	str := "Hello"
	fmt.Printf("Length is %d \n", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	for i, letters := range str {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letters))
	}

	fmt.Printf("Printing the character of string %c", str[0])
}
