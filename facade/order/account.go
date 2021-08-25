package main

import (
	"errors"
	"fmt"
)

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{name: accountName}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return errors.New("account name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}