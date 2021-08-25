package main

import "log"

func main() {
	walletFacade := newWalletFacade("qwer", 121234)
	err := walletFacade.addMoneyToWallet("qwer", 121234, 20)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	err = walletFacade.deductMoneyFromWallet("qwer", 121234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}