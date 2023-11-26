package main

type IBankAccount interface {
    GetBalance() int // 100 = 1 dollar
    Deposit(amount int)
    Withdraw(amount int) error
}
