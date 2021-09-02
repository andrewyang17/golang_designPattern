package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	b := &BankAccount{balance: balance}
	b.changes = append(b.changes, &Memento{balance})
	return b
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = $%d, currentPointer = %d", b.balance, b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	m := &Memento{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
	return m
}

func (b *BankAccount) Undo() {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
	}
}

func (b *BankAccount) Redo() {
	if b.current + 1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
	}
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)

	ba.Undo()
	fmt.Println("Undo 1", ba)
	ba.Undo()
	fmt.Println("Undo 2", ba)
	ba.Redo()
	fmt.Println("Redo", ba)
}