package main

import (
	"testing"
)

func TestWallet(t *testing.T){
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}

		//first deposit money in the wallet
		wallet.Deposit(Bitcoin(20))
		//withdraw the money
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want{
			t.Errorf("got %s want %s ", got, want)
		}
	})

}