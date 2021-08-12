package main

import "fmt"

type Employee struct {
	name, position string
	annualIncome int
}

func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			name:         name,
			position:     position,
			annualIncome: annualIncome,
		}
	}
}

func main() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Andrew")
	manager := managerFactory("Yang")
	fmt.Println(*developer)
	fmt.Println(*manager)
}
