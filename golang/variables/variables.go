package main

import "fmt"

func main() {
	a := true
	b1 := false

	fmt.Println("a && b1 ", a && b1)
	fmt.Println("a || b1 ", a || b1)

	var b, c int = 1, 2
	fmt.Println("b, c ", b, c)

	var k, h = 1, "hello"
	fmt.Println(k, h)

	var e int
	fmt.Println("e ", e)

	var g string
	fmt.Println(g)

	f := "apple"
	fmt.Println(f)

	// ep kieu du lieu
	i := 15
	j := 5.5
	sum := float64(i) + j /* sum := i + int(j) */
	println(sum)

	// khai báo nhiều biến
	var (
		firstname, lastname string
		age                 int
		salary              float64
		isConfirmed         bool
	)
	fmt.Println("firstname, lastname, age, salary, isConfirmed ", firstname, lastname, age, salary, isConfirmed)
}
