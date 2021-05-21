package main

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}
type Wallet struct {
	balance Bitcoin
}

//Deposit function
func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}

//Balance function
func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

//Withdraw function
func (w *Wallet) Withdraw(amount Bitcoin){
	w.balance -= amount
}