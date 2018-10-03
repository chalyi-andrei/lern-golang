package main

import (
	"fmt"
)

func main() {
	a := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range a {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
