package main

import "fmt"

func simpleClosure() func(a, b int) int {
	c := func(a, b int) int {
		return a + b
	}
	return c
}

func main() {
	c := simpleClosure()
	fmt.Println(c(3, 5))
}
