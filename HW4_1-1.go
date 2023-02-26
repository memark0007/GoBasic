package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mut sync.Mutex

var tobacco bool = false
var paper bool = false
var match bool = false

func smoker(name string, ready chan bool) {
	for {

		<-ready

		mut.Lock()

		if paper == true && match == true {

			paper = false
			match = false
			fmt.Println("Smoker 1 (tobacco) is smoking.")
		} else if tobacco == true && match == true {

			tobacco = false
			match = false
			fmt.Println("Smoker 2 (paper) is smoking.")
		} else if tobacco == true && paper == true {

			tobacco = false
			paper = false
			fmt.Println("Smoker 3 (match) is smoking.")
		}
		mut.Unlock()
	}

}

func agent(ready chan bool) {
	for {
		var ingred1, ingred2 int

		for {
			rand.Seed(time.Now().UnixNano())
			ingred1 = rand.Intn(3)
			ingred2 = rand.Intn(3)

			if ingred1 != ingred2 {
				break
			}
		}

		time.Sleep(time.Millisecond * 500)
		mut.Lock()
		if (ingred1 == 0 && ingred2 == 1) || (ingred1 == 1 && ingred2 == 0) && (tobacco != true && paper != true) {
			tobacco = true
			paper = true
			fmt.Println("Put tobacco and paper")
		} else if (ingred1 == 0 && ingred2 == 2) || (ingred1 == 2 && ingred2 == 0) && (tobacco != true && match != true) {
			tobacco = true
			match = true
			fmt.Println("Put tobacco and matches")
		} else if (ingred1 == 1 && ingred2 == 2) || (ingred1 == 2 && ingred2 == 1) && (paper != true && match != true) {
			paper = true
			match = true
			fmt.Println("Put paper and matches")
		}
		mut.Unlock()

		ready <- true

	}
}

func main() {
	ready := make(chan bool)
	go smoker("tobacco", ready)
	go smoker("paper", ready)
	go smoker("match", ready)
	go agent(ready)
	for {
	}
}
