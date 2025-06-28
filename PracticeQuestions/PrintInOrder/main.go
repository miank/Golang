package main

import (
	"fmt"
	"sync"
)

// func A(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	<-ch // wait for turn
// 	fmt.Println("A")
// 	ch <- 1 // signal next
// }

// func B(ch chan int, wg *sync.WaitGroup) {

// 	defer wg.Done()
// 	<-ch
// 	fmt.Println("B")
// 	ch <- 1
// }

// func C(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	time.Sleep(time.Millisecond)
// 	<-ch
// 	fmt.Println("C")
// }

func A1(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println("A")
		ch2 <- 1
	}
}

func B1(ch2, ch3 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-ch2
		fmt.Println("B")
		ch3 <- 1
	}
}

func C1(ch3, ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-ch3
		fmt.Println("C")
		ch1 <- 1
	}
}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(3)

	// go A(ch, &wg)
	// go B(ch, &wg)
	// go C(ch, &wg)

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	go A1(ch1, ch2, &wg)
	go B1(ch2, ch3, &wg)
	go C1(ch3, ch1, &wg)

	// Start the sequence
	ch1 <- 1 // allow A to proceed

	wg.Wait()
	close(ch)
}
