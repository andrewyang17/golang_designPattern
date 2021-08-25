package main

import "fmt"

const (
	TerroristDressType = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	// Flyweight: storing intrinsic state
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

// Flyweight factory
type dressFactory struct {
	dressMap map[string]dress
}

// To accept intrinsic state of the desired flyweight from a client
func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}