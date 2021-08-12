// Facade: Provides a simple, easy to understand/user interface over a large and
// sophisticated body of code. (You hide complicated system behind a very simple interface)

package main

import "fmt"

type Buffer struct {
	width, height int
	buffer []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: make([]rune, width*height),
	}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{
		buffer: buffer,
		offset: 0,
	}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

type Console struct {
	buffers []*Buffer
	viewports []*Viewport
	offset int
}

func NewConsole() *Console {
	b := NewBuffer(10, 10)
	v := NewViewport(b)
	return &Console{
		buffers:   []*Buffer{b},
		viewports: []*Viewport{v},
		offset:    0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func main() {
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}
//*