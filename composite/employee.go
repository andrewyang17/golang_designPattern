package main

import "fmt"

type Component interface {
	Display()
}

type Employee struct {
	id string
	name string
}

func (e Employee) Display() {
	fmt.Printf("ID: %s\n Name: %s\n", e.id, e.name)
}

func NewEmployee(id, name string) *Employee {
	return &Employee{id: id, name: name}
}

type Composite struct {
	component Component
	leaves []Component
}

func (c *Composite) Add(leaf Component) {
	c.leaves = append(c.leaves, leaf)
}

func (c *Composite) Display() {
	c.component.Display()
	if len(c.leaves) == 0 {
		return
	}
	fmt.Println("===List of Subordinates===")
	for _, leaf := range c.leaves {
		leaf.Display()
	}
	fmt.Println("===End===")
}

func NewComposite(component Component) *Composite {
	return &Composite{
		component: component,
	}
}

func ShowEmployeeHierarchy() {
	ceo := NewEmployee("ID-1", "Socrates")
	vpIdealist := NewEmployee("ID-2", "Plato")
	vpRealist := NewEmployee("ID-3", "Aristotle")
	directorIdealist := NewEmployee("ID-4", "Hegel")
	directorRealist := NewEmployee("ID-5", "Hume")

	directorIdealistOffice := NewComposite(directorIdealist)
	directorRealistOffice := NewComposite(directorRealist)

	vpIdealistOffice := NewComposite(vpIdealist)
	vpIdealistOffice.Add(directorIdealistOffice)

	vpRealistOffice := NewComposite(vpRealist)
	vpRealistOffice.Add(directorRealistOffice)

	company := NewComposite(ceo)
	company.Add(vpRealistOffice)
	company.Add(vpIdealistOffice)

	fmt.Println("=====Display Office of VP of Idealist=====")
	vpIdealistOffice.Display()
	fmt.Println("=====Display Office of VP of Realist=====")
	vpRealistOffice.Display()
	fmt.Println("=====Display Company=====")
	company.Display()
}

func main() {
	ShowEmployeeHierarchy()
}