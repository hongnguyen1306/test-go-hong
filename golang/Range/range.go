package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum ", sum)

	map1 := map[string]int{"k1": 1, "k2": 2}
	for k := range map1 {
		fmt.Println("key: ", k)
	}

	for i, val := range map1 {
		fmt.Println("i val ", i, val)
	}
}
