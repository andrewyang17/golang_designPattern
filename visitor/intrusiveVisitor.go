// Intrusive approach by definition a violation of the Open/Closed Principle

package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(*strings.Builder)
}

type DoubleExpression struct {
	value float64
}

func (d *DoubleExpression) Print(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (a *AdditionExpression) Print(builder *strings.Builder) {
	builder.WriteRune('(')
	a.left.Print(builder)
	builder.WriteRune('+')
	a.right.Print(builder)
	builder.WriteRune(')')
}

func main() {
	// 1 + (2+3)
	e := &AdditionExpression{
		left:  &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	// String Builder is the visitor, an intrusive builder because it modify the behavior
	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())
}
