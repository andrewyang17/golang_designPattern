package main

import "fmt"

const (
	developer = iota
	manager
)

type Employee3 struct {
	name, position string
	annualIncome int
}

func NewEmployee(role int, name string) *Employee3 {
	// Pre-configured object
	switch role {
	case developer:
		return &Employee3{name, "developer", 60000}

	case manager:
		return &Employee3{name, "manager", 80000}

	default:
		panic("unsupported role")
	}
}

func main() {
	developer := NewEmployee(developer, "Andrew Yang")
	developer.annualIncome = 65000
	fmt.Println(*developer)
}
