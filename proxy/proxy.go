// Proxy: A type that functions as an interface to a particular resources,
// and that resource may be remote, expensive to construct, or may require logging or
// some other added functionality.
// A proxy has the same interface as the underlying object.
// To create a proxy, simply replicate the existing interface of an object.
// Add relevant functionality to the redefined methods.

package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct {}

func (c *Car) Drive() {
	fmt.Println("Car being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{Age: 12})
	car.Drive()
}
