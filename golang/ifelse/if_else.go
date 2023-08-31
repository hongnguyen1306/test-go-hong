package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if a := 9; a < 8 {
		fmt.Println("a < 8")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("false")
	}
}
