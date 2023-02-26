// HW3_1_producer_consumer Channel Solution

 package main

 import (
	 "bytes"
	 "fmt"
	 "math/rand"
	 "sync"
	 "time"
 )
 
//  var m sync.Mutex
 var buffer = make([]byte, 0, 10)
 var all sync.WaitGroup
 
 func main() {
	 ch := make(chan []byte, 1)
	 ch <- buffer
	 rand.Seed(10)
	 all.Add(1)
	 go writer('a', ch)
	 go writer('b', ch)
	 go consumer(ch)
	 all.Wait()
 }
 
 func writer(c byte ,ch chan []byte) {
	 for i := 0; i < 5; i++ {
		 time.Sleep(time.Duration(rand.Int63n(1e9)))
		//  m.Lock()
		 lb := len(<- ch)
		//  lb := len(buffer)
		 if lb < cap(buffer) {
			 buffer = buffer[:lb+1]
			 buffer[lb] = c
			 fmt.Printf("'%c' written to buffer.     buffer contents: %s\n",
				 c, string(buffer))
		 }
		 ch <- buffer
		//  m.Unlock()
	 }
 }
 
 func consumer(ch chan []byte) {
	 a := []byte{'a'}
	 b := []byte{'b'}
	 for i := 0; i < 5; {
		 time.Sleep(time.Duration(rand.Int63n(1e9)))
		//  m.Lock()
		 bufferCh := <- ch
		 ai := bytes.Index(bufferCh, a)
		 bi := bytes.Index(bufferCh, b)
		 if ai >= 0 && bi >= 0 {
			 if ai > bi {
				 ai, bi = bi, ai
			 }
			 copy(buffer[bi:], buffer[bi+1:])
			 copy(buffer[ai:], buffer[ai+1:])
			 buffer = buffer[:len(buffer)-2]
			 fmt.Printf("pair removed from buffer.  buffer contents: %s\n",
				 string(buffer))
			 i++
		 }
		 ch <- buffer
		//  m.Unlock()
	 }
	 all.Done()
 }
 