/*
Viết chương trình nhập vào một danh sách các số (float) và tìm số lớn nhất, nhỏ nhất trong danh sách.
có handle lỗi nếu nhập vào không phải là 1 số
*/
package main

import (
	"fmt"
	"strconv"
)

func findMinMax(nums []float64) (float64, float64) {
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

func inputNumbers(n int) []float64 {
	numbers := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Nhập phần tử thứ %d: ", i+1)
		input := ""
		fmt.Scan(&input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Lỗi: Giá trị nhập vào không phải là số thực (float)")
			return nil
		}

		numbers[i] = num
	}

	return numbers
}

func mainTest() {
	var n int
	fmt.Print("Nhập số lượng phần tử (n): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Số lượng phần tử phải lớn hơn 0")
		return
	}

	numbers := inputNumbers(n)

	min, max := findMinMax(numbers)
	fmt.Println("Số bé nhất: ", min)
	fmt.Println("Số lớn nhất: ", max)
}
