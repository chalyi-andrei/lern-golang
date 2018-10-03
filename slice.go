package main

import (
	"fmt"
)

func main() {
	a := make([]string, 3)
	a[0] = "firstA"
	b := make([]string, len(a))
	copy(b, a)
	b[0] = "1"
	b[1] = "2"
	fmt.Println("a, b", a, b)
	fmt.Println("b[2]", b[2])
	b = append(b, "3")
	b = append(b, "4")
	b = append(b, "5")
	fmt.Println("b", b)

	x := b[4:5]
	fmt.Println('x', x)
}
