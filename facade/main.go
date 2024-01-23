package main

import "fmt"

func main() {

	wallet := Wallet{
		account:      &Account{account: 253142697809},
		security:     &Security{pin: 1991},
		balance:      &Balance{balance: 15000},
		notification: &Notification{},
	}
	wallet.addMoneyToWallet(253142697808, 1991, 1500)
	wallet.addMoneyToWallet(253142697809, 1990, 1500)
	wallet.addMoneyToWallet(253142697809, 1991, 1500)
	fmt.Println("----------------------------------------")
	wallet.debitMoneyFromWallet(253142697809, 1991, 15500)
	fmt.Println("----------------------------------------")
	wallet.debitMoneyFromWallet(253142697809, 1991, 1500)
}
