package main

import "fmt"

type Account struct {
	username string
	balance  int
}

func NewAccount(username string) *Account {
	return &Account{username: username}
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
	a.DisplayBalance()
}

func (a *Account) Withdraw(amount int) bool {
	if a.balance - amount > 0 {
		a.balance -= amount
		a.DisplayBalance()
		return true
	}
	return false
}

func (a *Account) DisplayBalance() {
	fmt.Printf("Hello %s, your account balance is %d\n", a.username, a.balance)
}
