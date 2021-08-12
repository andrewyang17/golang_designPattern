// builder: provides an API for constructing an object step-by-step (piecewise construction).
// Imagine having a factory function with 10 arguments (not productive).
// To make builder fluent, return the receiver -- allows chaining.

package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent+1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	b := HtmlBuilder{
		rootName: rootName,
		root:     HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
	return &b
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		name:    childName,
		text:     childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{
		name: childName,
		text: childText,
		elements: []HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)
	return b
}


func main() {

	example := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(example)
	sb.WriteString("</p>")
	fmt.Printf("%s\n", sb.String())

	sb.Reset()

	example2 := []string{"hello", "world"}
	for _, v := range example2 {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())

	example3 := NewHtmlBuilder("ul")
	example3.AddChild("li", "Home")
	example3.AddChild("li", "About")
	fmt.Println(example3.String())

	example4 := NewHtmlBuilder("select")
	example4.
		AddChildFluent("option", "Malaysia").
		AddChildFluent("option", "Singapore")
	fmt.Println(example4.String())
}
