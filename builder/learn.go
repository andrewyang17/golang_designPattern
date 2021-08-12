package main

import (
	"fmt"
	"strings"
)

type plan int
type payment int
type register func(n *NetflixRegistration)

const (
	mobile plan = iota
	basic
	standard
	premium

	card = iota
	postpaid
	code
)

type NetflixRegistration struct {
	email string
	plan plan
	payment payment
}

func (n *NetflixRegistration) EmailInput(email string) *NetflixRegistration {
	if !strings.Contains(email, "@") {
		panic("invalid email format")
	}
	n.email = email
	return n
}

func (n *NetflixRegistration) PlanInput(p plan) *NetflixRegistration {
	n.plan = p
	return n
}

func (n *NetflixRegistration) PaymentInput(p payment) *NetflixRegistration {
	n.payment = p
	return n
}

func SignUp(action register) {
	register := &NetflixRegistration{}
	action(register)
	fmt.Println("Account has successfully created!")
	fmt.Println("An account activation email has sent to", register.email)
}

func main() {
	SignUp(func(n *NetflixRegistration) {
		n.EmailInput("yangzhisiangg@gmail.com").
			PlanInput(mobile).
			PaymentInput(code)
	})
}
