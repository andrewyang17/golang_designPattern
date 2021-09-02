package main

import "fmt"

type freightTrain struct {
	mediator mediator
}

func (f *freightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		fmt.Println("freightTrain: Arrival blocked, waiting...")
		return
	}
	fmt.Println("freightTrain: Arrived")
}

func (f *freightTrain) depart() {
	fmt.Println("freightTrain: leaving...")
	f.mediator.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival() {
	fmt.Println("freightTrain: Arrival permitted, arriving...")
	f.arrive()
}
