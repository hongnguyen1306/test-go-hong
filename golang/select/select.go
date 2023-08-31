package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calA(c chan int) {
	randTime := rand.Intn(5)
	fmt.Printf("Chờ %v giây\n", randTime)
	time.Sleep(time.Duration(randTime) * time.Second)
	c <- 7
}

func calB(c chan int) {
	randTime := rand.Intn(5)
	fmt.Printf("Chờ %v giây\n", randTime)
	time.Sleep(time.Duration(randTime) * time.Second)
	c <- 8
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go calA(a)
	go calB(b)

	var selected int

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case resultA := <-a:
			selected = resultA
			fmt.Printf("Số được chọn là: %d", selected)
			return // nếu không có return thì chương trình sẽ chạy vòng lặp vô hạn
		case resultB := <-b:
			selected = resultB
			fmt.Printf("Số được chọn là: %d", selected)
			return
		default:
			fmt.Println("No value received")

		}
	}
}
