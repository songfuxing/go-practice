package main

import "fmt"

func mapTest() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("get ", v1)

	delete(m, "k1")
	fmt.Println("map ", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}
