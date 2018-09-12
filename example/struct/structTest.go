package main

import "fmt"

func main() {
	// new Struct
	fmt.Println(Person{"sfx", 12})
	fmt.Println(Person{name: "sfx"})

	// Access struct fields with a dot.
	s := Person{"sdd", 20}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.name)

	sp.name = "sfx"
	fmt.Println(s.name)

}

type Person struct {
	name string
	age  int
}
