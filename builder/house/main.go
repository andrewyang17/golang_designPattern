package main

import "fmt"

func main() {
	normal, err := getBuilder("normal")
	if err != nil {
		panic(err)
	}

	director := newDirector(normal)
	normalHouse := director.buildHouse()
	fmt.Println(normalHouse.Done())

	iglooBuilder, err := getBuilder("igloo")
	if err != nil {
		panic(err)
	}

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()
	fmt.Println(iglooHouse.Done())
}
