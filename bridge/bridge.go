// Bridge: A mechanism that decouples an interface (hierarchy)
// from an implementation (hierarchy).
// Connecting components together through abstractions.
// Prevents a 'Cartesian Product' complexity explosion
// A stronger form of encapsulation.

package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing lines for circle of radius", radius)
}

type RasterRenderer struct {}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	Renderer
	radius float32
}

func (c *Circle) Draw() {
	c.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{
		Renderer: renderer,
		radius:   radius,
	}
}

func main() {
	raster := RasterRenderer{}
	c1 := NewCircle(&raster, 5)
	c1.Draw()

	vector := VectorRenderer{}
	c2 := NewCircle(&vector, 5)
	c2.Resize(2)
	c2.Draw()
}
