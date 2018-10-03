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
	radius int
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := fmt.Println("perim:", r.perim())
}
