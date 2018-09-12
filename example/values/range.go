package main

import "fmt"

func rangeTest() {
	nums := []int{1, 2, 3, 4, 5}

	// range on arrays and slices provides both the index and value for each entry.
	for index, value := range nums {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "aaa", "b": "bbb"}
	for k, v := range kvs {
		fmt.Println("key:", k)
		fmt.Println("value:", v)
	}

	// range on strings iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
