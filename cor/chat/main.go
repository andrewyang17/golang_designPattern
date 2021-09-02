package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Flow interface {
	Execute(*bufio.Scanner, *visitor)
}

type visitor struct {
	query map[string]string
}

func NewVisitor() *visitor {
	return &visitor{query: make(map[string]string)}
}

func (v *visitor) GetJSON() []byte {
	bytes, _ := json.Marshal(v.query)
	return bytes
}

type handler struct {
	name     string
	question string
	next     Flow
}

func (h *handler) SetName(name string) {
	h.name = name
}

func (h *handler) SetQuestion(question string) {
	h.question = question
}

func (h *handler) SetNext(next Flow) {
	h.next = next
}

func (h *handler) Execute(scanner *bufio.Scanner, v *visitor) {
	fmt.Println(h.question)
	fmt.Print(">")
	scanner.Scan()
	v.query[h.question] = scanner.Text()
	fmt.Println()

	if h.next != nil {
		h.next.Execute(scanner, v)
	}
}

func NewHandler(name, question string) *handler {
	return &handler{
		name:     name,
		question: question,
	}
}

func main() {
	name := NewHandler("nameInput", "What is your name?")
	email := NewHandler("emailInput", "What is your email?")
	question := NewHandler("questionInput", "What is your question?")

	name.SetNext(email)
	email.SetNext(question)

	v := NewVisitor()

	scanner := bufio.NewScanner(os.Stdin)
	name.Execute(scanner, v)
	fmt.Println(string(v.GetJSON()))


}
