package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "first"
	c <- "second"
	close(c)

	for elem := range c {
		fmt.Println("elem", elem)
	}
}
