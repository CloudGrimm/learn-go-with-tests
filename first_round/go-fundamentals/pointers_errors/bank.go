package pointers_errors

import (
	"errors"
	"fmt"
)


type Bitcoin int

//structs
type Wallet struct {
	balance Bitcoin
}

//interfaces
type Stringer interface {
	String() string
}

//variables
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

/***functions***/
func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

//withdraw function
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance{
		return ErrInsufficientFunds
	}

	 w.balance -=amount
	 return nil
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC", b)
}