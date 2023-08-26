package main

import (
	"fmt"
)

func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

// func StrContains(arr []string, str string) bool {
// 	for _, a := range arr {
// 		if a == str {
// 			return true
// 		}
// 	}
// 	return false
// }

// func Unit64Contains(arr []uint64, str uint64) bool {
// 	for _, a := range arr {
// 		if a == str {
// 			return true
// 		}
// 	}
// 	return false
// }

func Contains[T comparable](arr []T, key T) bool { // comparable la kieu du lieu co the so sanh duoc
	for _, a := range arr {
		if a == key {
			return true
		}
	}
	return false
}

func createArr[K any](val ...K) []K {
	var arr []K
	arr = append(arr, val...)
	return arr
}

func main() {
	arr := []int{1, 2, 3}
	resultArr := Map(arr, func(v int) int { return v * 2 })
	fmt.Println(resultArr)

	arr2 := []string{"a", "b", "c"}
	resultArr2 := Map(arr2, func(v string) string { return v + v })
	fmt.Println(resultArr2)

	checkContains := Contains(arr, 1)
	if checkContains {
		fmt.Println("is Contain!!")
	} else {
		fmt.Println("not Contain!!!")
	}

	fmt.Println(createArr(1, 2, 3, 4, 5))
}
