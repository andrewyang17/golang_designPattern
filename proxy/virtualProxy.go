package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}

// Solution: virtual proxy
type LazyBitmap struct {
	filename string
	bitmap *Bitmap
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func main() {
	// Problem: Image doesn't get loaded until you actually need to render it
	fmt.Println("Not ideal")
	bmp := NewBitmap("demo.png")
	DrawImage(bmp)

	fmt.Println("Ideal")
	// When constructing NewLazyBitmap, it hasn't been materialized yet,
	// meaning the underlying implementation of bitmap hasn't even been constructed
	bmp2 := NewLazyBitmap("demo.png")
	// It's only been constructed whenever somebody explicitly ask for it
	DrawImage(bmp2)
}