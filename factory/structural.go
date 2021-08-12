package main

import "fmt"

type Factory interface {
	Create(name string) *Employee2
}

type EmployeeFactory struct {
	position string
	annualIncome int
}

type Employee2 struct {
	name string
	EmployeeFactory
}

func (f *EmployeeFactory) Create(name string) *Employee2 {
	return &Employee2{
		name:         name,
		EmployeeFactory: EmployeeFactory{
			position:     f.position,
			annualIncome: f.annualIncome,
		},
	}
}

func NewEmployeeFactory2(position string, annualIncome int) Factory {
	return &EmployeeFactory{
		position:    position,
		annualIncome: annualIncome,
	}
}

func main() {
	// Whoever is consuming this object has to know explicitly know that there is
	// a create method, we could overcome this by introducing interface
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	boss := bossFactory.Create("Steve Jobs")
	boss.annualIncome = 200000
	fmt.Println(*boss)
}
