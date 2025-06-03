package main

import (
	"fmt"
	"sync"
	"time"
)

func sendNumbers(nums []int, ch chan<- int) {
	for _, num := range nums {
		time.Sleep(time.Millisecond * 100)
		ch <- num
	}

	close(ch) // close the channel
}

func FanIn(ch1, ch2 <-chan int, out chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	readChannel := func(ch <-chan int) {
		for num := range ch {
			out <- num
		}
		wg.Done()
	}

	go readChannel(ch1)
	go readChannel(ch2)

	go func() {
		wg.Wait()
		close(out)
	}()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := make(chan int)

	go sendNumbers([]int{1, 2, 3}, ch1)
	go sendNumbers([]int{4, 5, 6}, ch2)

	go FanIn(ch1, ch2, out)

	for num := range out {
		fmt.Println("Numbers ", num)
	}

}
