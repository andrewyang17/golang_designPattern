package main

import "fmt"

type Age interface {Message()}
type underAge struct {age int}
type legalAge struct {age int}

func (*underAge) Message() {
	fmt.Println("You're underAge")
}

func (*legalAge) Message() {
	fmt.Println("Welcome!")
}

func NewCustomer(age int) Age {
	if age > 18 {
		return &legalAge{age: age}
	} else {
		return &underAge{age: age}
	}
}

func main() {
	c := NewCustomer(18)
	c.Message()
}
