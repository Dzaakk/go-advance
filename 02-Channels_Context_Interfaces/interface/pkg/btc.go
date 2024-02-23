package pkg

import "errors"

type BitCoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitCoinAccount {
	return &BitCoinAccount{
		balance: 0,
		fee:     150,
	}
}

func (b *BitCoinAccount) GetBalance() int {
	return b.balance
}

func (b *BitCoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitCoinAccount) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee

	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}
