package main

import (
	"fmt"
	"sync"
)

func A(chEven, chOdd chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		chEven <- 1
		if i%2 == 0 {
			fmt.Println(i)
		}
		<-chOdd
	}

}

func B(chOdd, chEven chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		chOdd <- 1
		if i%2 != 0 {
			fmt.Println(i)
		}
		<-chEven
	}
}

func main() {

	var wg sync.WaitGroup
	chEven := make(chan int, 1)
	chOdd := make(chan int, 1)

	wg.Add(2)
	go A(chEven, chOdd, &wg)
	go B(chOdd, chEven, &wg)

	chEven <- 1
	wg.Wait()

}
