// Timeout with Goroutine Response
// Problem Statement:
// Write a Go program where a goroutine performs a task that takes random time (1–3 seconds). If it completes within 2 seconds, print "Task completed". Otherwise, print "Timeout!".

package main

import (
	"fmt"
	"time"
)

func Task(done chan bool) {

	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {

		fmt.Println("Hello World")
	}

	done <- true
}

func main() {

	done := make(chan bool)

	go Task(done)

	select {
	case <-done:
		fmt.Println("Task Completed")
	case <-time.After(3 * time.Second):
		fmt.Println("TimeOut")

	}

}
