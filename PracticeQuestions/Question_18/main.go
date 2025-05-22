// Print Numbers in Order Using Multiple Goroutines
// Problem Statement:

// Write a Go program that launches three goroutines, each responsible for printing numbers in a specific range:

// Goroutine A prints numbers 1 to 3

// Goroutine B prints numbers 4 to 6

// Goroutine C prints numbers 7 to 9

// The numbers must be printed in exact order:
package main

import (
	"fmt"
	"sync"
)

func printA(start chan struct{}, next chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-start
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
	next <- struct{}{}
}

func printB(start chan struct{}, next chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-start
	for i := 4; i <= 6; i++ {
		fmt.Println(i)
	}
	next <- struct{}{}
}

func printC(start chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	<-start
	for i := 7; i <= 9; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	chA := make(chan struct{})
	chB := make(chan struct{})
	chC := make(chan struct{})

	go printA(chA, chB, &wg)
	go printB(chB, chC, &wg)
	go printC(chC, &wg)

	// Start the sequence
	chA <- struct{}{}

	wg.Wait()
}
