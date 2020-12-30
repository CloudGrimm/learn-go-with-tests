package main

import "testing"

func TestWallet(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)
		assertCorrectMessage(t, wallet, want)

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)
		want := Bitcoin(10)
		assertCorrectMessage(t, wallet, want)
	})
}
