package pointerserrors

import "fmt"

// Bitcoin of type int
type Bitcoin int

// Wallet struct to define data to be stored
type Wallet struct {
	balance Bitcoin
}

// Deposit into the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance for the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// String method to make use of stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
