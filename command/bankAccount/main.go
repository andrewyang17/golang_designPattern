package main

import "fmt"

func main() {
	// Single account
	ba := NewAccount("Andrew")

	cmdDeposit := NewAccountCommand(ba, Deposit, 100)
	cmdDeposit.Call()

	cmdWithdraw := NewAccountCommand(ba, Withdraw, 30)
	cmdWithdraw.Call()
	cmdWithdraw.Undo()

	// Transfer money
	fmt.Println("Transferring money...")
	ba2 := NewAccount("Yang")
	cmdTransfer := NewMoneyTransferCommand(ba, ba2, 40)
	cmdTransfer.Call()

	fmt.Println("Undoing...")
	cmdTransfer.Undo()
}