package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println
	rand.Seed(int64(time.Now().Nanosecond())) // for uniq number every time

	i := rand.Intn(100)
	p(i)
}
