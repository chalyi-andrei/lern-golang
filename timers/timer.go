package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getData() int {
	rand.Seed(int64(time.Now().Nanosecond()))
	r := rand.Intn(10)
	return r
}

func main() {
	tiker := time.NewTicker(time.Second)

	go func() {

		for t := range tiker.C {
			data := getData()

			switch {
			case data < 5:
				fmt.Println("data < 5", t)
			case data > 5:
				fmt.Println("data > 5, tiker is stop!", t)
				tiker.Stop()
			}

		}
	}()

	time.Sleep(time.Second * 6)
	tiker.Stop()

	fmt.Println("stop")
}
