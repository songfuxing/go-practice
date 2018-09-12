package main

import "fmt"

func main() {
	i := 1
	fmt.Println("initial:", i)
	toZeroVal(i)
	fmt.Println("val:", i)
	toZeroPoint(&i)
	fmt.Println("point:", i)

	fmt.Println("address:", &i)
}

// 传值
func toZeroVal(v int) {
	v = 0
}

// 传指针
func toZeroPoint(p *int) {
	*p = 0
}
