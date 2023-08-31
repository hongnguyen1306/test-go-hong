package main

import "fmt"

func main() {
	// var s []string
	// fmt.Println("uninit ", s, s == nil, len(s) == 0)

	// s = make([]string, 3)
	// fmt.Println(s, "cap(s) ", cap(s), "len(s)", len(s))

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("s ", s)
	// Một Slice sẽ có 2 thuộc tính là length (len) và capacity (cap). Length là số phần tử chứa trong Slice, còn capacity là số phần tử chứa trong Array
	//  mà Slice tham chiếu đến (bắt đầu tính từ phần tử đầu tiên của Slice). Để lấy ra length của Slice ta dùng hàm len(), còn để lấy ra capacity thì ta dùng hàm cap().
	arr := []int{1, 2, 3, 4, 5}
	// slice := arr[0:1]
	// //fmt.Println("slice, len(slice), cap(slice) ", slice, len(slice), cap(slice))
	// fmt.Println(slice)

	// slice[0] = 4

	//fmt.Println(slice, "arr", arr)

	// arr1 := make([]int, 0)
	// arr1 = append(arr1, arr[0])

	// arr1[0] = 5

	// fmt.Println(arr1, "arr", arr)

	// tham chieu
	arr1 := arr
	arr1[0] = 55
	fmt.Println(arr1, "arr", arr) // arr1 và arr thay đổi [55 2 3 4 5] arr [55 2 3 4 5]

	// tham tri toi arr
	arrTT := [...]int{1, 2, 3, 4, 5}
	arr2 := arrTT
	arr2[0] = 100
	fmt.Println(arr2, "arrTT", arrTT) // [100 2 3 4 5] arrTT [1 2 3 4 5]

}
