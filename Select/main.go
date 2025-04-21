package main

import (
	"fmt"
	"time"
)

// If you must use default, only do it for non-blocking checks.
// Otherwise, it’ll skip even if channels aren’t ready.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from Channel1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from Channel2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

}
