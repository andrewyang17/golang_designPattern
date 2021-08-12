package main

import "fmt"

type Worker interface {
	DoWork()
}

type Company struct {
	Employee []Worker
}

func NewCompany(workers ...Worker) *Company {
	c := &Company{}
	for _, worker := range workers {
		c.Employee = append(c.Employee, worker)
	}
	return c
}

func (c *Company) GetWorkers() []Worker {
	return c.Employee
}

type Designer struct {}

func (*Designer) DoWork() {
	fmt.Println("Designer is designing...")
}

func NewDesigner() *Designer {
	return &Designer{}
}

type Architect struct {}

func (*Architect) DoWork() {
	fmt.Println("Architect is drawing...")
}

func NewArchitect() *Architect {
	return &Architect{}
}

func main() {
		d := NewDesigner()
		a := NewArchitect()

		c := NewCompany(d, a)

		for _, worker := range c.GetWorkers() {
			worker.DoWork()
		}
}