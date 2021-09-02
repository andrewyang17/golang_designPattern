package main

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type AccountCommand struct {
	account   *Account
	action    Action
	amount    int
	succeeded bool
}

func (a *AccountCommand) Call() {
	switch a.action {
	case Deposit:
		a.account.Deposit(a.amount)
	case Withdraw:
		a.succeeded = a.account.Withdraw(a.amount)
	}
}

func (a *AccountCommand) Undo() {
	if !a.succeeded {return}
	switch a.action {
	case Deposit:
		a.account.Withdraw(a.amount)
	case Withdraw:
		a.account.Deposit(a.amount)
	}
}

func (a *AccountCommand) SetSucceeded(value bool) {
	a.succeeded = value
}

func (a *AccountCommand) Succeeded() bool {
	return a.succeeded
}

func NewAccountCommand(account *Account, action Action, amount int) *AccountCommand {
	return &AccountCommand{account: account, action: action, amount: amount}
}