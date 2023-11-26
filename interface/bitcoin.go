package main

import "errors"

type Bitcoin struct {
    balance int
    fee int
}

func NewBitcoin() *Bitcoin {
    return &Bitcoin{
	balance: 0,
	fee: 300,
    }
}

func (b *Bitcoin) GetBalance() int {
    return b.balance
}

func (b *Bitcoin) Deposit(amount int) {
    b.balance += amount
}

func (b *Bitcoin) Withdraw(amount int) error {
    newBalance := b.balance - amount - b.fee
    if newBalance < 0 {
	return errors.New("insufficient funds")
    }
    b.balance = newBalance
    return nil
}

