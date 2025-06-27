package main

import (
	"fmt"
	"sync"
)

func Ping(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("*** Ping ***")
		ch <- 1
	}

}

func Pong(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		ch <- 1
		fmt.Println("^^^ Pong ^^^")
		<-ch
	}
}

func main() {

	ch := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(2)

	go Ping(ch, &wg)
	go Pong(ch, &wg)

	//	ch <- 1

	wg.Wait()
	close(ch)

}
