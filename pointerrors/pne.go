package pointerrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin float64

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

// Wallet
// Pproperties: balance Bitcoin
// Wallet that holds a balance of Bitcoin
type Wallet struct {
	balance Bitcoin
}

// Accepts: ammount Bitcoin
// Deposit an ammount of Boitcoin in Wallet
func (w *Wallet) Deposit(ammount Bitcoin) {
	w.balance += ammount
}

// Returns Bitcoin
// Check and return the balance of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Accpets: ammount Bitcoin
// Returns: error ErrInsufficientFunds
// Withdraw an ammount from Wallet and error if insufficient funds
func (w *Wallet) Withdraw(ammount Bitcoin) error {

	if ammount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= ammount
	return nil
}
