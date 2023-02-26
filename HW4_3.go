package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numberOfBees = 6
	potSize      = 10
)

var (
	mutex  sync.Mutex
	cond   sync.Cond
	honey  int
	buffer [potSize]int
	front  int
	rear   int
	isFull bool
	prod   [numberOfBees]*sync.WaitGroup
	wg     *sync.WaitGroup
)

func putHoney(value int) {
	buffer[front] = value
	front = (front + 1) % potSize
}

func eatHoney() int {
	tmp := buffer[rear]
	rear = (rear + 1) % potSize
	return tmp
}

func bee(no int) {
	for {
		mutex.Lock()
		for isFull {
			mutex.Unlock()
			cond.Wait()
			mutex.Lock()
		}

		position := front
		putHoney(honey)
		fmt.Println("The bee No.", no, "has put honey to the pot")
		honey++

		if position == potSize-1 {
			isFull = true
		}
		mutex.Unlock()
		time.Sleep(2 * time.Second)
	}
}

func bear() {
	for {

		if isFull {
			mutex.Lock()

			for i := 0; i < potSize; i++ {
				eatHoney()
			}
			fmt.Println("Bear is finished eating honey")
			isFull = false
			cond.Signal()
			mutex.Unlock()
		}
	}
}

func main() {
	cond.L = &mutex
	wg = &sync.WaitGroup{}
	wg.Add(1)
	go bear()
	for i := 0; i < numberOfBees; i++ {
		prod[i] = &sync.WaitGroup{}
		prod[i].Add(1)
		go bee(i)
	}

	for i := 0; i < numberOfBees; i++ {
		prod[i].Wait()
	}
	wg.Wait()
}
