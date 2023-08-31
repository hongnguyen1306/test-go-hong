package main

import "fmt"

func main() {
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		fmt.Println(n)
		return fib(n-1) + fib(n-3) // 5 4 3 2
	}

	// fmt.Println(fib(5))
	fib(6)
}
