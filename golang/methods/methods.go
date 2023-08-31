package main

import "fmt"

type Person struct {
	name, city  string
	age, salary int
}

func (p Person) displayPerson() {
	fmt.Printf("Salary of %s is %d at %d in %s", p.name, p.salary, p.age, p.city)
}

func (p *Person) changeAge(newAge int) {
	p.age = newAge
}

type rectangle struct {
	width  int
	length int
}

func area(r rectangle) int {
	return r.length * r.width
}

func (r rectangle) area() int {
	return r.length * r.width
}

func main() {
	person1 := Person{
		name:   "Lee",
		city:   "HCM",
		age:    33,
		salary: 100000,
	}
	person1.displayPerson()
	fmt.Println(person1)

	person1.changeAge(44)
	fmt.Println(person1)

	rec1 := rectangle{
		width:  5,
		length: 10,
	}
	fmt.Println(rec1.area())
	fmt.Println(area(rec1))

	// Pointer
	rec2 := &rec1
	fmt.Println(rec2.area())
	fmt.Println(area(*rec2))
}
