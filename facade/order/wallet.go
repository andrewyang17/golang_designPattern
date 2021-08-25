package main

import (
	"errors"
	"fmt"
)

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) debitBalance(amount int) {
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	fmt.Printf("Your new balance is %d\n", w.balance)
}

func (w *wallet) creditBalance(amount int) error {
	if w.balance < amount {
		return errors.New("balance is not sufficient")
	}
	w.balance = w.balance - amount
	return nil
}