package main

import "fmt"

func main() {

	a := 6
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	if a > 3 && a < 6 {
		fmt.Println("a is within range")
	}

	c := 1
	b := 2

	if c > b {
		fmt.Println("c is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}

	// For loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break
		}
	}

	// continue
	for i := i; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Nested for loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Outer loop iteration %d\n", i)
		for j := 0; j < 2; j++ {
			fmt.Printf("i = %d j = %d \n", i, j)
		}
	}

	// There is not while loop .... same can be achieved using for loop

	k := 1
	for k <= 5 {
		fmt.Println(k)
		k++
	}

}
