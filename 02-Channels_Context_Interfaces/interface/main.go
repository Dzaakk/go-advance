package main

import (
	"fmt"
	"interface/pkg"
)

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	myAccounts := []IBankAccount{
		pkg.NewWellsFargo(),
		pkg.NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdraw(100); err != nil {
			panic(err)
		}

		balance := account.GetBalance()
		fmt.Printf("balacnce = %d\n", balance)
	}
}
