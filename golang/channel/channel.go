// Dung trong mo hinh cac tac vu chay dong thoi Concurrency Programming
// Goroutine su dung channel de tuong tac voi nhau
// Chia se gia tri thong qua channel

/* 	Unbuffered channel yeu cau ca 2 goroutines gửi và nhận sẵn sàng cùng một lúc.
	+ a := make(chan int)
	+ Hàm main là một goroutine.
	+ Lúc gửi dữ liệu thì goroutine bị block để gửi dữ liệu vào channel. Vậy nên cần có một goroutine lấy dữ liệu
để unblock.
*/

/*	Buffered channel có khoảng trống để chứa dữ liệu, ko cần cả 2 goroutines phải sẵn sàng gửi và nhận cùng một lúc.
	+ a := make(chan int, 10)
	+ Thường được dùng khi cần sự bất đồng bộ để xử lý cho nhanh.
*/

package main

import "fmt"

// pipeline
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, a := range nums {
			out <- a
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	fmt.Println("in ", <-in)
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func receiveAndSend(c chan int) {
	fmt.Printf("Received: %d\n", <-c)
	fmt.Printf("Sending 2...\n")
	c <- 2
}

func receiveOnly(c <-chan int) { // Hàm chỉ nhận dữ liệu
	fmt.Printf("Recevied: %d\n", <-c)
	// c <- 2  Error khi cố tình gửi dữ liệu
}

func main() {

	messages := make(chan string)

	/* Đoạn code dưới đây sẽ bị deadlock vì ở dòng 21, main goroutine bị block để gửi dữ liệu vậy nên các dòng tiếp theo
	   không được thực thi */
	// messages <- "hello"
	// go func() {
	// 	println(<-messages)
	// }()

	// Cách 1:
	go func() { messages <- "hello" }()
	fmt.Println(<-messages) // OK

	// Cách 2:
	go func() {
		fmt.Println(<-messages)
	}()
	messages <- "hi" // OK

	// check close channel
	out := gen(2, 5)
	// for {
	// 	value, isAlive := <-out
	// 	if !isAlive {
	// 		fmt.Printf("Value: %d. Channel has been closed.\n", value)
	// 		break
	// 	}
	// 	fmt.Printf("Value: %d\n", value)
	// }

	result := sq(out)
	fmt.Printf("Result : %d\n", <-result)
	fmt.Printf("Result : %d\n", <-result)

	myChan := make(chan int)
	go receiveAndSend(myChan)
	myChan <- 1
	fmt.Printf("Value from receiveAndSend: %d\n", <-myChan)

	// For Range
	valueChan := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			valueChan <- i
		}
		close(valueChan)
	}()

	for a := range valueChan {
		fmt.Printf("Value: %d\n", a)
	}
	go receiveOnly(valueChan)

}
