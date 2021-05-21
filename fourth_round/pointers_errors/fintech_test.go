package main

import (
	"testing"
)

func TestWallet(t *testing.T){

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertCorrectMessage(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertCorrectMessage(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)

		//if err == nil{
		//	t.Errorf("wanted an error but didn't get one")
		//}
	})

}

func assertCorrectMessage(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError (t testing.TB, err error, want error) {
	t.Helper()
	if err ==nil {
		t.Fatal("wanted an error but did not get one")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
