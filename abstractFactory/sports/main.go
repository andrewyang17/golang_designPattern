//You need to buy a sports kit, a set of two different products: a pair of shoes and a shirt.
// You would want to buy a full sports kit of the same brand to match all the items.

package main

import (
	"fmt"
)

func PrintShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func PrintShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	adidasShirt := adidasFactory.getShirt()
	adidasShoe := adidasFactory.getShoe()

	nikeShirt := nikeFactory.getShirt()
	nikeShoe := nikeFactory.getShoe()

	PrintShirtDetails(adidasShirt)
	PrintShoeDetails(adidasShoe)

	PrintShirtDetails(nikeShirt)
	PrintShoeDetails(nikeShoe)
}