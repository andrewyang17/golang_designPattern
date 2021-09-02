package main

type rectangle struct {
	Height int
	Width  int
}

func newRectangle(height int, width int) shape {
	return &rectangle{Height: height, Width: width}
}

func (t *rectangle) accept(v visitor) {
	v.visitForRectangle(t)
}

func (t *rectangle) getType() string {
	return "Rectangle"
}
