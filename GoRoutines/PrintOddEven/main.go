package main

import (
	"fmt"
	"sync"
)

func Even(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			<-ch // Wait for signal
			fmt.Printf("Even Element %d\n", i)
			ch <- true
		}
	}

}

func Odd(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 != 0 {
			<-ch // Wait for signal
			fmt.Printf("Odd Element %d\n", i)
			ch <- true
		}
	}
}

func main() {
	ch := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go Even(ch, &wg)
	go Odd(ch, &wg)
	ch <- true // Start with true to print even
	wg.Wait()
	close(ch)
}
