package main

import (
	"fmt"
	"time"
)

func Aname() {
	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}
	for t1 := 0; t1 < 3; t1++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

func Aid() {
	arr2 := [4]int{300, 301, 302, 303}
	for t2 := 0; t2 <= 3; t2++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {
	go display("Welcome")

	display("GeeksforGeeks")

	// Anonymous Goroutine

	fmt.Println("Welcome !! to main function")
	go func() {
		fmt.Println("Welcome!! to GeeksforGeeks")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Good Bye!! to Main function")

	// Multiple GoRoutine
	fmt.Println("!....... Main Go-routine Start ....!")
	go Aname()

	go Aid()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n ...Main Go-routine End...!")
}
