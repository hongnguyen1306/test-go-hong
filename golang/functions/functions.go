package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func mulReturns(a, b int) (int, int) {
	return a + b, a * b
}

func findX(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums) // %T\n type of ...
	found := false
	for _, val := range nums {
		if val == num {
			found = true
		}
	}
	if found {
		fmt.Println("Found X")
	} else {
		fmt.Println("Not Found X")
	}
}

func changeString(s ...string) {
	s[0] = "Change"
	s = append(s, "Life")
	fmt.Println(s)
}

func main() {

	// Functions
	sum := plus(3, 4)
	fmt.Println("sum ", sum)

	// Multiple Return Values
	result1, result2 := mulReturns(3, 4)
	fmt.Println(result1, result2)

	// Variadic Functions (Hàm bất định)
	findX(2, 1, 3, 4, 5)

	nums := []int{3, 4, 5, 6}
	findX(1, nums...)

	s1 := []string{"Hello Go"}
	changeString(s1...)
	fmt.Println(s1)

}
