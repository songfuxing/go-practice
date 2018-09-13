package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 0 {
		return -1, errors.New("can not work with &arg")
	}
	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("ck-%d,%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 0 {
		return -1, &argError{arg, "can not work with it"}
	}
	return arg, nil
}

func main() {
	for _, i := range []int{0, 2} {
		if arg, e := f1(i); e != nil {
			fmt.Println("fail:", e)
		} else {
			fmt.Println("work:", arg)
		}
	}

	for _, i := range []int{0, 2} {
		if arg, e := f2(i); e != nil {
			fmt.Println("fail:", e)
		} else {
			fmt.Println("work:", arg)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
