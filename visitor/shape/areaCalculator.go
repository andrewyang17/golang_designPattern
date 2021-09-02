package main

import (
	"fmt"
	"math"
)

type areaCalculator struct {}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("The area is", s.Side * s.Side)
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("The area is",  3.142 * math.Pow(float64(c.Radius), 2))
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	fmt.Println("The area is", r.Height * r.Width)
}
