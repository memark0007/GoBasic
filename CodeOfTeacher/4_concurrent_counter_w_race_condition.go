package main

import (
	"fmt"
	"sync" 
	"runtime"
)

var wg sync.WaitGroup
var sum = 0

func getGOMAXPROCS() int {
    return runtime.GOMAXPROCS(0)
}

func counter(n int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		sum = sum + 1 

	}
	fmt.Println("From ", n, ":", sum)
	wg.Done()

}

func main() {
	
    fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
