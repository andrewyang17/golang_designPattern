package main

type circle struct {
	Radius int
}

func newCircle(radius int) shape {
	return &circle{Radius: radius}
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}