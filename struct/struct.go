package main

import "fmt"

type Address struct {
	state, city string
}

type person struct {
	name string
	age  int
}

type school struct {
	string
	int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 45
	return &p
}

func main() {
	p := person{
		name: "Ann",
		age:  23,
	}

	fmt.Println(p)
	fmt.Println(person{name: "Ross"})
	fmt.Println(&person{name: "Finn"})

	p2 := newPerson("Lee")
	fmt.Println("p2 ", p2)
	p2.age = 100
	fmt.Println("p4 ", p2)

	// s := school{"123School", 4}
	// fmt.Println(s)

	// // So sanh struct
	// person1 := person{name: "Ann", age: 12}
	// person2 := person{name: "Lii", age: 12}
	// fmt.Println("== ", person1 == person2) // so sanh gia tri & cau truc

	// person3 := &person1
	// person3.age = 13
	// fmt.Println(*person3, person1)

}
