package main

import (
	"fmt"
)

type kg float32

const (
	groundCostPerKg = 1.5
	airCostPerKg = 3.0
)

type ShippingMethod interface {
	getShippingCost(w ShippingWeight) float32
}

type ShippingWeight interface {
	getWeight() kg
}

type Product struct {
	Name string
	Price float32
	Quantity int
	Weight kg
}

func NewProduct(name string, price float32, quantity int, weight kg) *Product {
	return &Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Weight:   weight,
	}
}

type Order struct {
	Orders []Product
	ShippingMethod
}

func NewOrder(method ShippingMethod, products ...*Product) *Order {
	var p []Product
	for _, product := range products {
		p = append(p, *product)
	}
	return &Order{
		Orders:         p,
		ShippingMethod: method,
	}
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, order := range o.Orders {
		totalPrice += order.Price * float32(order.Quantity)
	}
	return round(totalPrice)
}

func (o *Order) TotalPayment() float32 {
	return round(o.TotalPrice() + o.getShippingCost(o))
}

func (o *Order) getWeight() kg {
	var totalWeight kg
	for _, order := range o.Orders {
		totalWeight += order.Weight
	}
	return totalWeight
}

type Ground struct {}

func (*Ground) getShippingCost(w ShippingWeight) float32 {
	return float32(w.getWeight()) * groundCostPerKg
}

type Air struct {}

func (*Air) getShippingCost(w ShippingWeight) float32 {
	return float32(w.getWeight()) * airCostPerKg
}

func round(v float32) float32 {
	return float32(int(v*100)) / 100
}

func main() {
	p := NewProduct("Salmon", 12.9, 2, 2)
	p2 := NewProduct("Seaweed", 6, 2, 0.5)
	p3 := NewProduct("Sushi Rice", 9.9, 1, 1)

	g := Ground{}
	o := NewOrder(&g, p, p2, p3)
	fmt.Println("Total before shipping cost is ", o.TotalPrice())
	fmt.Println("Total after shipping cost is ", o.TotalPayment())

	a := Air{}
	o2 := NewOrder(&a, p, p2, p3)
	fmt.Println("Total before shipping cost is ", o2.TotalPrice())
	fmt.Println("Total after shipping cost is ", o2.TotalPayment())
}