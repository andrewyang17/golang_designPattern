// Decorator: Facilitates the addition of behaviors to individuals objects
// through embedding.
// Do not want to rewrite or alter existing code (OCP)
// Want to keep new functionality separate (SRP)
// A decorator embeds the decorated object.
// Adds utility fields and methods to augment the object's features.
// Often used to emulate multiple inheritance (can be problem if you embed two objects
// which have same field, you have to end up having to manage the consistency across those
// different objects).

package main

import "fmt"

type Shape interface {
	Render() string
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius %2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square wide side %2f", s.Side)
}

func main() {
	circle := Circle{2}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())
}
