package main

import (
	"fmt"
	"log"
)

func main() {
	vm := newVendingMachine(1, 10)

	err := vm.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vm.addItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vm.requestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.insertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}