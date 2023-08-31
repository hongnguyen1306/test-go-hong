package main

import "fmt"

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("This program is packing with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {
	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("HELP") // panic here

}

func main() {
	fmt.Println("main() started")

	normMain() // normal call

	fmt.Println("main() done")
}
