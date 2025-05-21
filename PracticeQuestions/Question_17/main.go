//  Print "Ping" and "Pong" Alternately Using Goroutines

package main

import (
	"fmt"
	"time"
)

func Ping(ch chan int) {
	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("Ping")
		ch <- 1
	}

}

func Pong(ch chan int) {
	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("Pong")
		ch <- 1
	}
}

func main() {

	ch := make(chan int)

	go Ping(ch)
	go Pong(ch)

	ch <- 1
	time.Sleep(1 * time.Second)
}
