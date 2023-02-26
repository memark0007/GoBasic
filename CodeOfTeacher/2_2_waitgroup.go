package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup) {

	fmt.Println("hello from goroutine!")
	wg.Done() //notify complete task to the group
}

func main() {
	fmt.Println("Hello, Go hackers!")

	// Create a wait group.
	var wg sync.WaitGroup

	wg.Add(1) //add # of goroutine to wait group
	go hello(&wg)
	wg.Add(1)
	go hello(&wg)

	// Wait until everyone finishes.
	wg.Wait()
	fmt.Println("Done")

}
