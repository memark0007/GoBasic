package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main(){

	for i := 0; i < 5; i++ {
		fmt.Println("Main: hello")
		wg.Add(1)
		go routine(&wg)
		wg.Wait()
	}

}
func routine(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("Goroutine: world")
}