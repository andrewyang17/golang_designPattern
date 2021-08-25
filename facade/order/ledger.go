package main

import "fmt"

type ledger struct {}

func (s *ledger) makeEntry(accountID, taxType string, amount int) {
	fmt.Printf("Make ledger entry for accountID %s with txnType %s for amount %d\n",
		accountID, taxType, amount)
}
