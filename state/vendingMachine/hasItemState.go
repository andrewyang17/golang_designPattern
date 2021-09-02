package main

import (
	"errors"
	"fmt"
)

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h *hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return errors.New("no item present")
	}
	fmt.Println("Item requested")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *hasItemState) insertMoney(money int) error {
	return errors.New("please select item first")
}

func (h *hasItemState) dispenseItem() error {
	return errors.New("please select item first")
}

