package main

import "fmt"

type Person2 struct {
	name, position string
}

type personMod func(*Person2)

type PersonBuilder2 struct {
	actions []personMod
}

func (b *PersonBuilder2) Called(name string) *PersonBuilder2 {
	b.actions = append(b.actions, func(p *Person2) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder2) WorkAsA(position string) *PersonBuilder2 {
	b.actions = append(b.actions, func(p *Person2) {
		p.position = position
	})
	return b
}

func (b *PersonBuilder2) Build() *Person2 {
	p := Person2{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

func main() {
	b := PersonBuilder2{}
	p := b.
		Called("Andrew Yang").
		WorkAsA("Programmer").
		Build()
	fmt.Println(p)
}