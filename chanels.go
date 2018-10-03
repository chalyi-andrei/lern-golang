package main

import (
	"fmt"
	"time"
)

func worker(channel chan bool) {
	fmt.Print("start...")
	time.Sleep(time.Second)
	fmt.Println("done")

	channel <- true
}

func main() {
	// buffering
	str := make(chan string, 2)

	str <- "Hello"
	str <- "world!"

	fmt.Println("str", <-str)
	fmt.Println("str", <-str)


	// Synchronization
	someChannel := make(chan bool, 1)

	go worker(someChannel)
	<-someChannel

	// Timeout

	c := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c <- "result"
	}()

	select {
	case res:= <-c:
		fmt.Println(res)
	case <-time.After(time.Second * 3): // if we set 2 seconds, timeout will works
		fmt.Println("timeout!")
	}
}