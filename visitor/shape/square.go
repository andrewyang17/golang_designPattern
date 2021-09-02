package main

type square struct {
	Side int
}

func newSquare(side int) shape {
	return &square{Side: side}
}

func (s *square) getType() string {
	return "Square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}
