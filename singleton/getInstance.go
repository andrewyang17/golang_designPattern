package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var once2 sync.Once
var singleInstance *single

type single struct {}

func getInstance(wg *sync.WaitGroup) *single {
	defer wg.Done()

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// nil-check after the lock is acquired.
		// Ensure that if more than one goroutine bypasses the first check,
		// only one goroutine can create the singleton instance.
		if singleInstance == nil {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}


// Using sync.Once
func getInstance2(wg *sync.WaitGroup) *single {
	defer wg.Done()

	if singleInstance == nil {
		once2.Do(func() {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created")
	}
	return singleInstance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getInstance2(&wg)
	}
	wg.Wait()
}
