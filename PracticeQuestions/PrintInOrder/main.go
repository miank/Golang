package main

import (
	"fmt"
	"sync"
	"time"
)

func A(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch // wait for turn
	fmt.Println("A")
	ch <- 1 // signal next
}

func B(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	<-ch
	fmt.Println("B")
	ch <- 1
}

func C(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	<-ch
	fmt.Println("C")
}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(3)

	go A(ch, &wg)
	go B(ch, &wg)
	go C(ch, &wg)

	// Start the sequence
	ch <- 1 // allow A to proceed

	wg.Wait()
	close(ch)
}
