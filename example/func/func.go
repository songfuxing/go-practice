package main

import "fmt"

func funcTest() {
	res := plus(1, 2)
	fmt.Println(res)
	result := plusPlus(1, 2, 3)
	fmt.Println(result)

	fmt.Println(multi(1))
	_, right := multi(1)
	fmt.Println(right)

	variadic(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	variadic(nums...)

	nextInt := closures()
	fmt.Println(nextInt(), nextInt(), nextInt())
	// new index
	newInt := closures()
	fmt.Println(newInt())

	fmt.Println(recursion(3))
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// multi return value
func multi(a int) (int, int) {
	return a, a + 1
}

func variadic(nums ...int) {
	fmt.Println(nums)
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func recursion(a int) int {
	if a == 1 {
		return a
	}
	return a * recursion(a-1)
}
