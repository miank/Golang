package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once sync.Once

type single struct {
}

var (
	singleInstance,
	singleInstance_1 *single
)

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created 1.")
		}
	} else {
		fmt.Println("Single instance already created 2.")
	}

	return singleInstance
}

func getInstance_1() *single {
	if singleInstance_1 == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now with Once")
				singleInstance_1 = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance_1
}

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}

	for i := 0; i < 10; i++ {
		go getInstance_1()
	}
	fmt.Scanln()
}
