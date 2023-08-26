/*
Viết chương trình nhập vào một danh sách các số (float) và tìm số lớn nhất, nhỏ nhất trong danh sách.
có handle lỗi nếu nhập vào không phải là 1 số
*/
package main

import (
	"fmt"
	"reflect"
)

func findMinMax(nums ...float32) (float32, float32) {
	for _, num := range nums {
		result := fmt.Sprintf("%s", reflect.TypeOf(num))
		if result != "float32" {
			panic("value is not a float")
		}
	}
	var (
		min = nums[0]
		max = nums[0]
	)
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return min, max
}

func main() {
	fmt.Println(findMinMax(6, 3, 4, 5, 6))
}
