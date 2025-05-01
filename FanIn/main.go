package main

import (
	"fmt"
	"sync"
)

func someStrings() <-chan string {
	ch := make(chan string)
	numbersString := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
	}

	go func() {
		for _, numberString := range numbersString {
			ch <- numberString
		}
		close(ch)
	}()

	return ch
}

func someNumbers() <-chan string {
	ch := make(chan string)
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	go func() {
		for _, number := range numbers {
			ch <- number
		}

		close(ch)

	}()

	return ch
}

func fanIn(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup

	multiplexStream := make(chan string)

	multiplex := func(c <-chan string) {
		defer wg.Done()

		for i := range c {
			multiplexStream <- i
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexStream)
	}()

	return multiplexStream

}

func main() {
	ch1 := someNumbers()
	ch2 := someStrings()

	mergeCh := fanIn(ch1, ch2)

	go func() {
		for val := range mergeCh {
			fmt.Println(val)
		}
	}()

}
