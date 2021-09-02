package main

import (
	"errors"
)

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) addItem(count int) error {
	n.vendingMachine.incrementItemCount(count)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *noItemState) requestItem() error {
	return errors.New("item out of stock")
}

func (n *noItemState) insertMoney(money int) error {
	return errors.New("item out of stock")
}

func (n *noItemState) dispenseItem() error {
	return errors.New("item out of stock")
}
