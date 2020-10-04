package pointerserrors

// Wallet struct to define data to be stored
type Wallet struct {
	balance int
}

// Deposit into the wallet
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance for the wallet
func (w *Wallet) Balance() int {
	return w.balance
}
