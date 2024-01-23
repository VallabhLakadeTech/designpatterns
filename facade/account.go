package main

import (
	"errors"
)

type Account struct {
	account int
}

func (acc Account) verifyAccount(account int) error {
	if acc.account != account {
		return errors.New("Account not matched. Please try again.")
	}
	return nil
}
