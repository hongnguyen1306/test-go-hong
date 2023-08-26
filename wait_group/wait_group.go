// Wait Group thường được dùng để chờ các goroutines con chạy xong (đếm goroutines ở phần wg.Add(2))
// đến khi nào về 0 thì main goroutine mới được chạy.

/*
	Race condition: Tương tranh, các tiến trình cùng truy cập vào một nguồn tài nguyên và cùng thay đổi chúng.
	 				Việc truy cập ko tuần tự dẫn đến sai lệch kết quả so với mong muốn.
	Lúc này cần dùng đến Channel, Atomic và Mutex để các goroutines đi theo tuần tự.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

func test1(wg *sync.WaitGroup) {
	fmt.Println("fUNC 1")
	wg.Done()
}

func test2(wg *sync.WaitGroup) {
	fmt.Println("FUNC 2")
	wg.Done()
}

func cal(wg *sync.WaitGroup) {
	atomic.AddInt64(&x, 1) // DÙNG ATOMIC để xử lý đồng bộ, xử lý dữ liệu kiểu nguyên int32, uint64, uint32,...
	wg.Done()
}

func main() {
	var wg, w sync.WaitGroup
	wg.Add(2)
	go test1(&wg)
	go test2(&wg)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go cal(&w)
	}

	wg.Wait()
	w.Wait()
	fmt.Println("Giá trị của X là: ", x)
	fmt.Println("Doneeee")
}
