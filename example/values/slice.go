package main

import "fmt"

func slice() {
	s := make([]string, 5)
	fmt.Println("empty:", s)
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"
	s[3] = "4"
	s[4] = "5"
	fmt.Println("set:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c=", c)
	fmt.Println("c[2:5]=", c[2:5])
	fmt.Println("c[0:1]=", c[0:1])
	fmt.Println("c[:2]=", c[:2])
}
