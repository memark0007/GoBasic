package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func barber(customers chan int, customerReady chan bool, barberReady chan bool, freeChairs chan bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for {
		select {
		//wait for customer
		case <-customerReady:
			mutex.Lock()
			customerNumber := <-customers
			fmt.Println("Barber: Cutting hair of customer", customerNumber)
			time.Sleep(time.Duration(rand.Intn(500)+1) * time.Millisecond)
			mutex.Unlock()
			fmt.Println("Barber: Customer", customerNumber, "is done")
			fmt.Println("Free Chairs (wating room):", 5-len(freeChairs))
			barberReady <- true
		//wait for new customer
		case <-time.After(time.Second * 3):
			wg.Done()
		}
	}
}

func customer(i int, customers chan int, barberReady chan bool, customerReady chan bool, freeChairs chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Customer", i, "arrived")
	select {
	case freeChairs <- true:
		customers <- i
		fmt.Println("Customer", i, "takes a seat")
		fmt.Println("Free Chairs (wating room):", 5-len(freeChairs))
		//is barber ready?.
		<-barberReady
		fmt.Println("Customer", i, "getting hair cut")
		//customer is ready.
		customerReady <- true
		<-freeChairs
		fmt.Println("Free Chairs (wating room):", 5-len(freeChairs))
	default:
		fmt.Println("***Customer", i, "leaves the shop")
	}

}
func main() {
	customers := make(chan int, 15)
	barberReady := make(chan bool, 1)
	customerReady := make(chan bool, 1)
	freeChairs := make(chan bool, 5)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	barberReady <- true
	wg.Add(1)
	go barber(customers, customerReady, barberReady, freeChairs, &mutex, &wg)
	rand.Seed(10)
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		time.Sleep(time.Duration(rand.Intn(100)+1) * time.Millisecond)
		go customer(i, customers, barberReady, customerReady, freeChairs, &wg)
	}
	wg.Wait()
	fmt.Println("No more customers coming.")
	fmt.Println("Closing shop.")

}
