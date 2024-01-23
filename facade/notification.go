package main

import "fmt"

type Notification struct {
}

func (notification Notification) sendCreditNotification(amount int) {
	fmt.Printf("Transaction completed. %v amount credited to your account.\n", amount)
}

func (notification Notification) sendDebitNotification(amount int) {
	fmt.Printf("Transaction completed. %v amount debitted from your account.\n", amount)
}
