package main

import (
	"fmt"
	"sync"
	"time"
)

func DoJob(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	ch <- 1

	fmt.Printf("Job %d started \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Job %d done \n", id)
	<-ch
}

func main() {

	N := 10
	K := 3

	var wg sync.WaitGroup

	ch := make(chan int, K)

	for i := 1; i <= N; i++ {
		wg.Add(1)
		go DoJob(i, ch, &wg)
	}

	wg.Wait()
	fmt.Println("All jobs completed")

}
