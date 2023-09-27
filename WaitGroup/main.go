package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	var waitGroup sync.WaitGroup
	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go concurrentTask(i, &waitGroup)
	}

	waitGroup.Wait()
	finishTask()
	fmt.Printf("Program end \n")
}

func concurrentTask(taskNumber int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)

}

func finishTask() {
	fmt.Printf("Executing finish task")
}
