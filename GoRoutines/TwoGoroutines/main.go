package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of main Goroutine")
	go greetings("John")
	go greetings("Mary")
	time.Sleep(time.Second * 10)
	fmt.Println("End of main Goroutine")

}

func greetings(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, "==>", name)
		time.Sleep(time.Millisecond)
	}
}
