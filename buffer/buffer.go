/*
	Buffered channel là một channel trong golang có khả năng lưu trữ dữ liệu bên trong nó.
	Cú pháp khai báo: chanName := make(chan Type, capacity)
	Ví dụ: chanName := make(chan int, 5)
*/
/*
	+ BufferedChan không block main goroutine khi nó còn sức chứa, vậy nên không cần thiết phải lấy dữ liệu từ channel ngay
		Bufferhan has len = 0, cap = 5
		Bufferhan has len = 1, cap = 5
		Bufferhan has len = 2, cap = 5
		Bufferhan has len = 1, cap = 5
	+ Lấy dữ liệu từ bufferedChan trống cũng sẽ làm goroutine bị block
	+ Lưu trữ dữ liệu theo thứ tự FIFO (First In First Out)
*/

package main

import (
	"fmt"
	"time"
)

const (
	numberOfUrls    = 10000
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
}

func startQueue() <-chan int {

	queue := make(chan int, 100)
	go func() {
		for i := 1; i <= numberOfUrls; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		close(queue)
	}()
	return queue
}

func main() {

	// Giải quyết bài toán lấy 10K URLs với buffered channel và 5 goroutines
	queue := startQueue()

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, fmt.Sprintf("%d", i))
	}
	time.Sleep(time.Minute * 2)
	// end!

	// Gửi dữ liệu vượt quá 5 nhưng k bị deadlock do goroutine
	// ở trong go func, nếu vượt quá sẽ bị khoá và chờ dữ liệu đc lấy ra sau đó gửi tiếp
	bufferedChan := make(chan int, 5)
	go func() {
		fmt.Println("For StartQueue!!!!")
		for i := 1; i <= 10; i++ {
			bufferedChan <- i
			fmt.Printf("bufferedChan %d\n", i)
		}
		close(bufferedChan)
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(<-bufferedChan)
	}
	//end!!

	//FIFO (First In First Out)
	newChan := make(chan int, 10)
	for i := 1; i < 10; i++ {
		newChan <- i
	}
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(<-newChan)
	// }

	// Khác biệt với Unbufferd Channel
	bufChan := make(chan int, 1)
	// unbufChan := make(chan int)

	bufChan <- 1 // OK
	/* unbufChan <- 1  Deadlock */

	// Crawl 10k URLs với buffered channel và goroutines
}
