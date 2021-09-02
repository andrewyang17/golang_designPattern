package main

func main() {
	stationManager := NewStationManager()

	pt := &passengerTrain{mediator: stationManager}
	ft := &freightTrain{mediator: stationManager}

	pt.arrive()
	ft.arrive()
	pt.depart()
}