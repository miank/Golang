package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	// Producer goroutine
	go func() {
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("Produced: ", i)
		}
		close(ch)
	}()

	// Consumer goroutine
	go func() {
		for num := range ch {
			fmt.Println("Consumed: ", num)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(3 * time.Second)
}
