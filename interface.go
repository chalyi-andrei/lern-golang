package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

// Rectangle
type rect struct {
	width, height int
}
func (r *rect) area() int {
	return r.width * r.height
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Circle
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{width: 3, height: 2}
	fmt.Println("perim:", r.perim())
	fmt.Println("area:", r.area())
	fmt.Println("r:", r)
}
