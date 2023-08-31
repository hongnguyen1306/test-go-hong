package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func calc(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Can't not work with 0")
	}
	return a / b, nil
}

func main() {
	result, err := calc(6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result: ", result)
	}
}
