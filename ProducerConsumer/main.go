package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("producing items ", i)
		ch <- i
		time.Sleep(time.Second)
	}
}

func consumer(ch chan int) {
	for item := range ch {
		fmt.Println("Consuming item", item)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int, 2)
	go producer(ch)
	go consumer(ch)
	time.Sleep(10 * time.Second)
}
