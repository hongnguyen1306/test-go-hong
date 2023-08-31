package main

import (
	"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Constract struct {
	empID    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Constract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.CalculateSalary()
		fmt.Printf("Total Expense Per Month %d \n", expense)
	}
}

func main() {
	pemp1 := Permanent{1, 3000, 2}
	pemp2 := Permanent{2, 5000, 4}
	cemp1 := Constract{3, 4000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
