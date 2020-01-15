package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is just an int
type Bitcoin int

// implements stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a wallet
type Wallet struct {
	balance Bitcoin
}

// Balance gets the balance in the wallet
func (w *Wallet) Balance() Bitcoin {
	// return (*w).balance
	// is the same as
	return w.balance
}

// Deposit will deposit
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds is an error for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw withdraws
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
