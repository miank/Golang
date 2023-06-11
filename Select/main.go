package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	// Print both channel using loop

	go goThree(ch3)
	go goFour(ch4)

	for i := 0; i < 2; i++ {
		select {
		case msg3 := <-ch3:
			fmt.Println(msg3)
		case msg4 := <-ch4:
			fmt.Println(msg4)
		}
	}
}

func goOne(ch chan string) {
	//	time.Sleep(time.Millisecond * 1)

	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	// time.Sleep(time.Millisecond * 1)

	ch <- "From goTwo goRoutine"
}

func goThree(ch chan string) {
	//	time.Sleep(time.Millisecond * 1)

	ch <- "From goThree goroutine"
}

func goFour(ch chan string) {
	// time.Sleep(time.Millisecond * 1)

	ch <- "From goFour goRoutine"
}
