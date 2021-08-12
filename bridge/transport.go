package main

import "fmt"

type Mobile int

const (
	Truck Mobile = iota
	Car
)

type Transporter interface {
	Move()
	Navigate()
}

type Transport struct {
	Transporter
}

func (t *Transport) Deliver(destination, from string) {
	t.Move()
	t.Navigate()
	fmt.Printf("The items is delivering from %v to %v \n", from, destination)
}

func NewTransport(t Transporter) *Transport {
	return &Transport{t}
}

type Tesla struct {
	autonomous bool
	eco        bool
	types      Mobile
}

func (t *Tesla) Move() {
	if t.eco {
		fmt.Println("Car battery is fully charged")
	} else {
		fmt.Println("Car fuel is fully filled")
	}
}

func (t *Tesla) Navigate() {
	if t.autonomous {
		fmt.Println("Robot is taking over")
	} else {
		fmt.Println("A human driver is assigned")
	}
}

func NewTesla(t Mobile, a, e bool) *Tesla {
	return &Tesla{a, e, t}
}


func main() {
	t := NewTesla(Car, false, false)
	t2 := NewTesla(Truck, true, true)

	n := NewTransport(t)
	n2 := NewTransport(t2)
	n.Deliver("Kuala Lumpur", "Penang")
	n2.Deliver("Penang", "Kuala Lumpur")
}
