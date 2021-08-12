// Prototype: A partially or fully initialized object that you copy and make use of.
// Reiterate existing designs.
// An existing design is a prototype, we deep copy the prototype.

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite int
	StreetAddress, City string
}

type Employee struct {
	Name string
	Office *Address
}

func (e *Employee) DeepCopy() *Employee {
	// With serialization
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	_ = encoder.Encode(e)

	d := gob.NewDecoder(&b)
	result := &Employee{}
	_ = d.Decode(result)

	return result
	// Without using serialization
	//return &Employee{
	//	Name:  e.Name,
	//	Office: &Address{
	//		Suite:         e.Office.Suite,
	//		StreetAddress: e.Office.StreetAddress,
	//		City:          e.Office.City,
	//	},
	//}
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	newE := proto.DeepCopy()
	newE.Name = name
	newE.Office.Suite = suite

	return newE
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewSubOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&subOffice, name, suite)
}

var mainOffice = Employee{
	Name:   "Steve Jobs",
	Office: &Address{
		Suite:         1,
		StreetAddress: "Waverley Street",
		City:          "Palo Alto",
	},
}

var subOffice = Employee{
	Name:   "Sherlock Holmes",
	Office: &Address{
		Suite:         2,
		StreetAddress: "221B Baker Street",
		City:          "London",
	},
}

func main() {
	boss := NewMainOfficeEmployee("Masayoshi son", 3)
	boss2 := NewSubOfficeEmployee("Lebron James", 4)
	fmt.Println(mainOffice, mainOffice.Office)
	fmt.Println(subOffice, subOffice.Office)
	fmt.Println(boss, boss.Office)
	fmt.Println(boss2, boss2.Office)
}
