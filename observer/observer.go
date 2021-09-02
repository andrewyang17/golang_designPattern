package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

//type PropertyChanged struct {
//	Name  string
//	Value interface{}
//}

type Person struct {
	Observable
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Observable: Observable{subs: new(list.List)},
		name:       name,
		age:        age,
	}
}

func (p *Person) CatchACold() {
	p.Fire(p.name)
}

//func (p *Person) GetAge() int {
//	return p.age
//}
//
//func (p *Person) SetAge(age int) {
//	if age == p.age {return}
//	p.age = age
//	p.Fire(PropertyChanged{"Age", p.age})
//}
//
//type TrafficManagement struct {
//	o Observable
//}
//
//func (t *TrafficManagement) Notify(data interface{}) {
//	if pc, ok := data.(PropertyChanged); ok {
//		if pc.Value.(int) >= 18 {
//			fmt.Println("Congrats, you can drive now!")
//			t.o.Unsubscribe(t)
//		}
//	}
//}

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s\n", data.(string))
}

func main() {
	p := NewPerson("AndrewYang", 13)
	ds := &DoctorService{}
	ds2 := &DoctorService{}
	p.Subscribe(ds)
	p.Subscribe(ds2)

	p.CatchACold()

	//t := &TrafficManagement{p.Observable}
	//p.Subscribe(t)
	//
	//for i := 14; i <= 20; i++ {
	//	fmt.Println("Setting the age to", i)
	//	p.SetAge(i)
	//}
}
