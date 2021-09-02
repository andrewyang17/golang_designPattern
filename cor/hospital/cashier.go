package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment done")
		c.next.execute(p)
		return
	}
	fmt.Println("Cashier getting money from patient")
	p.paymentDone = true
}

func (c *cashier) setNext(next department) {
	c.next = next
}
