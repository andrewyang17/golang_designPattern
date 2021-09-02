package main

type MoneyTransferCommand struct {
	CompositeAccountCommand
	from, to *Account
	amount int
}

func NewMoneyTransferCommand(from *Account, to *Account, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{
		from:                    from,
		to:                      to,
		amount:                  amount,
	}
	c.commands = append(c.commands, NewAccountCommand(from, Withdraw, amount))
	c.commands = append(c.commands, NewAccountCommand(to, Deposit, amount))
	return c
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}