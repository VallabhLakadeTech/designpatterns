package main

import (
	"errors"
	"fmt"
)

type Balance struct {
	balance int
}

func (acc Balance) credit(amount int) {
	acc.balance += amount
	fmt.Printf("Amount credited: %v. Total balance: %v.\n", amount, acc.balance)
}

func (acc Balance) debit(amount int) error {
	if acc.balance < amount {
		return errors.New("Insufficient balance.")
	}
	acc.balance -= amount
	fmt.Printf("Amount debitted: %v. Total balance: %v.\n", amount, acc.balance)
	return nil
}
