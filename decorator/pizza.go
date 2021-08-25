package main

import "fmt"

type pizza interface {
	getPrice() int
}

type pizzahut struct {}

func (p *pizzahut) getPrice() int {
	return 10
}

// decorators
type extraCheese struct {
	pizza pizza
}

func (ec *extraCheese) getPrice() int {
	pizzaPrice := ec.pizza.getPrice()
	return pizzaPrice + 3
}

type extraTopping struct {
	pizza pizza
}

func (et *extraTopping) getPrice() int {
	pizzaPrice := et.pizza.getPrice()
	return pizzaPrice + 3
}

func main() {
	pizza := &pizzahut{}

	pizzaExtraCheese := &extraCheese{pizza}
	fmt.Println("Price with extra cheese:", pizzaExtraCheese.getPrice())

	pizzaExtraTopping := &extraTopping{pizzaExtraCheese}
	fmt.Println("Price with extra topping:", pizzaExtraTopping.getPrice())
}