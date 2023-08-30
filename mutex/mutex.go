package main

import (
	"fmt"
	"sync"
)

var x int64 = 0
var mutex = &sync.Mutex{}

func addOne(wg *sync.WaitGroup, i int) {
	fmt.Printf("into goroutine %d\n", i)
	// Lock lại
	mutex.Lock()
	x = x + 1
	// Unlock
	mutex.Unlock()

	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 20; i++ {
		w.Add(1)
		fmt.Printf("goroutine %d\n", i)
		go addOne(&w, i)
	}
	w.Wait()
	fmt.Println("Giá trị của x là: ", x)
}
