// ** La phuong thuc chay dong thoi cac ham/ phuong thuc khac. gon nhe, chi phi thap

package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello() // co them tu "go" truoc ham, KHONG cho gorotine chay xong, chay tiep dong code tiep sau
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
