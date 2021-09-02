package main

import "fmt"

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting...")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *passengerTrain) depart() {
	fmt.Println("PassengerTrain: leaving...")
	p.mediator.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving...")
	p.arrive()
}
