package main

import (
	"fmt"
	"sync"
)
// var mu sync.Mutex
var wg sync.WaitGroup
var sum = 0

func counter(n int, wg *sync.WaitGroup, ch chan int){
	for i := 0; i < 10000; i++ {
		// mu.Lock()
		sumIncounter:= <- ch
		sum = sumIncounter + 1
		ch <- sum
		// mu.Unlock()
	}
	fmt.Println("from ", n, ":", sum)
	wg.Done()
}

func main(){
	ch := make(chan int,1)	
	ch <- 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg, ch)
	}
	wg.Wait()
	sum = <- ch
	fmt.Println("Final Sum:", sum)
}
