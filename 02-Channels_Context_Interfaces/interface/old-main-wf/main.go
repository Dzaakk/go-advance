package main

import (
	"fmt"
	"interface/pkg"
)

type BankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	wf := pkg.NewWellsFargo()

	wf.Deposit(200)
	wf.Deposit(200)
	if err := wf.Withdraw(100); err != nil {
		panic(err)
	}

	currentBalance := wf.GetBalance()

	fmt.Printf("WF balance: %d", currentBalance)
}
