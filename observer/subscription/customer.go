package main

import "fmt"

type customer struct {
	email string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.email, itemName)
}

func (c *customer) getID() string {
	return c.email
}

func newCustomer(email string) *customer {
	return &customer{email: email}
}

