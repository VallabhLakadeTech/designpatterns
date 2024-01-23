package main

import "errors"

type Security struct {
	pin int
}

func (security Security) verifyPin(pin int) error {
	if security.pin != pin {
		return errors.New("Invalid pin. Please try again.")
	}
	return nil
}
