package main

import (
	"errors"
	"fmt"
)

type MyGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

func newAk47() MyGun {
	return &ak47{gun{
		name:  "AK47",
		power: 4,
	}}
}

type musket struct {
	gun
}

func newMusket() MyGun {
	return &musket{gun{
		name:  "Musket",
		power: 2,
	}}
}

func PrintDetails(g MyGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

// Factory
func CreateGun(gunType string) (MyGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	} else {
		return nil, errors.New("wrong gun type passed")
	}
}

func main() {
	w, _ := CreateGun("ak47")
	w2, _ := CreateGun("musket")

	PrintDetails(w)
	PrintDetails(w2)
}