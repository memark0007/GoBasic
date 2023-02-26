package main

import (
	"fmt"
	"time"
)

func main(){
	
	ch := make(chan string,1)	
	ch <- "ping"
	go Player("Alice", ch)
	go Player("Bob", ch)

	time.Sleep(time.Second * 3)
}

func Player(name string,c chan string)  {
	for i := 0; i < 5; i++ {
		x := <-c
		if x == "ping"{			
			fmt.Println(name+": ping")
			c <- "pong"
		}else if x == "pong"{
			fmt.Println(name+": pong")
			c <- "ping"
		}
	}
}
