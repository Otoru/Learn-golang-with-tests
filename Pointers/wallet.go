package main

import (
	"errors"
	"fmt"
)

// Bitcoin represents a number of coins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// The Wallet stores the number of coins someone currently has
type Wallet struct {
	balance Bitcoin
}

// Deposit add Bitcoins to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds means that the wallet does not have enough Bitcoin for the requested transaction
var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

// Withdraw subtracts some Bitcoin from the wallet
// Return an error if it cannot be performed
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns total of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {}
