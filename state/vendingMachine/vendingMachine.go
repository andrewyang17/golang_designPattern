package main

import "fmt"

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState  state

	itemCount     int
	itemPrice     int
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{itemCount: itemCount, itemPrice: itemPrice}
	v.hasItem = &hasItemState{vendingMachine: v}
	v.itemRequested = &itemRequestedState{vendingMachine: v}
	v.hasMoney = &hasMoneyState{vendingMachine: v}
	v.noItem = &noItemState{vendingMachine: v}

	v.setState(v.hasItem)

	return v
}


