// Question: Write a Go program with two goroutines that send messages to a single channel.
// One goroutine sends "Hello" three times, and the other sends "World" three times.
// Use a simple fan-in pattern to collect these messages into one channel and print them in the main function.
// Stop the program after all messages are received, and make sure the channel is closed properly.

// Message: Hello
// Message: World
// Message: Hello
// Message: World
// Message: Hello
// Message: World
package main

import (
	"fmt"
	"sync"
	"time"
)

func HelloWorld(ch chan string, msg string, cnt int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < cnt; i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- msg
	}
}

func main() {
	ch := make(chan string, 6)
	var wg sync.WaitGroup

	wg.Add(2)
	go HelloWorld(ch, "Hello", 3, &wg)
	go HelloWorld(ch, "World", 3, &wg)

	// Close the channel once both goroutines are done - two ways to wait.
	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	wg.Wait()
	close(ch)

	// Receive from channel until it's closed
	for msg := range ch {
		fmt.Println("Message:", msg)
	}
}
