package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a ", a)

	var b = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(b)

	for index, value := range a {
		value = b[index] + 1
		fmt.Println(index, value)
	}
}
