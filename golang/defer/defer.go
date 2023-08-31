package main

import "fmt"

func main() {
	var err error
	defer func() {
		fmt.Println(err)
		fmt.Println(222)
	}()

	defer func() {
		fmt.Println(333)
	}()

	defer func() {
		fmt.Println(444)
	}()

	err = A()
	if err != nil {
		return
	}

	err = B()
	if err != nil {
		return
	}

	err = C()
	if err != nil {
		return
	}
}

func A() error {
	//return fmt.Errorf("A")
	return nil
}

func B() error {
	//return fmt.Errorf("B")
	return nil
}

func C() error {
	return fmt.Errorf("C")
}
