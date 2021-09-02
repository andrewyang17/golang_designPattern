package main

func main() {
	s := newSquare(5)
	c := newCircle(7)
	r := newRectangle(10 ,5)

	areaCalculator := &areaCalculator{}

	s.accept(areaCalculator)
	c.accept(areaCalculator)
	r.accept(areaCalculator)
}