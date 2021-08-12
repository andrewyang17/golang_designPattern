// Factory: A component responsible solely for the wholesale (not piecewise) creation of objects.
// It is a helper function for making struct instances.
// It is any entity that can take care of object creation.

package main

import "fmt"

type Person struct {
	Name string
	Age int
	EyeCount int
}

// Factory function: some additional logic or default value needs to be applied
// when initializing a particular object
func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("Andrew Yang", 25)
	fmt.Println(p)
}
