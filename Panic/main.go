package main

import "fmt"

func entry(lang *string, aname *string) {

	// Defer statement
	defer fmt.Println("Defer statement is the entry function \n")

	if lang == nil {
		panic("Error : Language cannot be nil")
	}

	if aname == nil {
		panic("Error : Author name cannot be nil")
	}

	fmt.Printf("Author Langauge: %s \n Author Name: %s\n", *lang, *aname)

}

func main() {
	A_Lang := "Go Language"

	defer fmt.Printf("Defer statement in the Main function \n")
	entry(&A_Lang, nil)

}
