package main

import (
	"fmt"
)

func main2() {
	// array
	nums := [3]float64{1.23, 2.534, 3.6453}
	fmt.Println("array", nums)
	for i, value := range nums {
		fmt.Println(value, i)
	}

	// map
	users := make(map[string]string)
	users["proger"] = "Vasya"
	users["teacher"] = "Misha"
	fmt.Println("map", users)

	// functions
	var a = 2
	var b = 3
	var res1, res2 = summ(a, b)
	fmt.Println("function ", res1, res2)

	//pointer
	var x = 0
	pointer(&x)
	fmt.Println("pointer:", x)

	// struct (object)
	testCat := Cats{"Bob", 5, 0.67}
	fmt.Println("Bob age is", testCat.age)
}

func summ(a int, b int) (int, string) {
	var res int
	res = a + b

	return res, "some string"
}

func pointer(x *int) {
	*x = 2
}

type Cats struct {
	name     string
	age      int
	happines float64
}
