// Let’s say there are 5 Terrorists and 5 Counter-Terrorist, so a total of 10 players.
// Now there are two options concerning dress.
// 1. Each of the 10 player objects creates a different dress object and embeds them.
// A total of 10 dress objects will be created.
// 2. We create two dress objects
// Single Terrorist Dress Object: This will be shared across 5 Terrorists.
// Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist.


package main

import "fmt"

func main() {
	game := newGame()

	//Add Terrorist
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	//Add CounterTerrorist
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}