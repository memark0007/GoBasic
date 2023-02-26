package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(i int, cs chan string) {
	i = i + 1
	cakeName := "Strawberry Cake " + strconv.Itoa(i)
	fmt.Println("Making a cake and sending ...", cakeName)
	cs <- cakeName //send a strawberry cake
}
func receiveCakeAndPack(cs chan string) {
	s := <-cs //get whatever cake is on the channel
	fmt.Println("Packing received cake: ", s)
}

func main() {

	cs := make(chan string)

	for i := 0; i < 3; i++ {

		go makeCakeAndSend(i, cs)
		go receiveCakeAndPack(cs)

		//sleep for a while so that the program doesn’t exit immediately and output is clear for illustration
		time.Sleep(1 * 1e9)
	}
}
