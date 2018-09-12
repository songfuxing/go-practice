package main

import "fmt"

func main() {
	r := rect{2, 3}
	fmt.Println(r.aere())
	fmt.Println(r.perimeter())
}

type rect struct {
	width, height int
}

func (r *rect) aere() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2 * (r.width * r.height)
}
