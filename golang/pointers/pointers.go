package main

import "fmt"

// func change(val *int) {
// 	*val = 55
// }

// func modify(arr *[3]int) {
// 	arr[0] = 200 //or (*arr)[0] = 200
// }

// func main() {
// 	b := 255
// 	a := &b
// 	fmt.Println("address b ", a)
// 	fmt.Println("value b ", *a)
// 	*a = 123
// 	fmt.Println("value *a, b ", *a, b)

// 	// Pointer trong hàm
// 	change(a)
// 	fmt.Println("value of a before function call is", *a)

// 	// Pointer trong slice
// 	arr := [3]int{1, 2, 3}
// 	modify(&arr)
// 	fmt.Println(arr)

// }

func main() {
	names := []string{"Stanley", "David", "Oscar"}

	persons := make([]Person, 0)

	for _, name := range names {
		n := name
		fmt.Println(&n)
		persons = append(persons, Person{Name: &n})
	}

	for _, p := range persons {
		fmt.Println(*p.Name)
	}
}

type Person struct {
	Name *string
}
