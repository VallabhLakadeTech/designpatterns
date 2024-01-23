package main

import "fmt"

type Wallet struct {
	account      *Account
	balance      *Balance
	security     *Security
	notification *Notification
}

func (wallet Wallet) addMoneyToWallet(account, security, amount int) {
	err := wallet.account.verifyAccount(account)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = wallet.security.verifyPin(security)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wallet.balance.credit(amount)
	wallet.notification.sendCreditNotification(amount)
}

func (wallet Wallet) debitMoneyFromWallet(account, security, amount int) {
	err := wallet.account.verifyAccount(account)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = wallet.security.verifyPin(security)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = wallet.balance.debit(amount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wallet.notification.sendDebitNotification(amount)
}
