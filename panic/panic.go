package main

import (
	"fmt"
	"log"
)

func checkPanic(a []int, index int) {
	if index > len(a)-1 {
		panic("Out of bound access foe slice") // gọi panic tường minh
	}
	fmt.Println(a[index])
}

func panic(interface{})
func recover() interface{}

func main() {
	// a := []int{2, 3}
	// // fmt.Println(a[2]) // vượt quá giới hạn mảng

	// // checkPanic(a, 2)

	// defer fmt.Println("Defer in main")
	// fmt.Println("After painc in f2 ", a)
	// panic("Panic with Defer")

	if r := recover(); r != nil {
		log.Fatal(r)
	}

	panic(123)

	if r := recover(); r != nil {
		log.Fatal(r)
	}

}
