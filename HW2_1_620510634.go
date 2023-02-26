package main

import (
	"fmt"
	"sync"
)
var mu sync.Mutex
var wg sync.WaitGroup
var sum = 0

func counter(n int, wg *sync.WaitGroup){
	for i := 0; i < 10000; i++ {
		mu.Lock()
		sum = sum + 1
		mu.Unlock()
	}
	fmt.Println("from ", n, ":", sum)
	wg.Done()
}

func main(){
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}
	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
