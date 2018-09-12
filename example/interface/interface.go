package main

import (
	"fmt"
	"math"
)

type geometer interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometer) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
func main() {
	r := rect{2, 3}
	c := circle{5}

	measure(r)
	measure(c)
}
