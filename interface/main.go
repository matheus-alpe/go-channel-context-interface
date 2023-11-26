package main

import "fmt"

func main() {
    myAccounts := []IBankAccount{
	NewWellsFargo(),
	NewBitcoin(),
    }

    for _, account := range myAccounts {
	fmt.Printf("Account = %T\n", account)
	account.Deposit(500)

	if err := account.Withdraw(260); err != nil {
	    fmt.Printf("ERR = %v\n", err)
	}

	balance := account.GetBalance()
	fmt.Printf("balance = %d\n", balance)
    }
}
