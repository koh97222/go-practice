package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want BitCoin) {

		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			// そもそもnilで帰ってきてしまう場合、それ以降のアサーションに進むのを防ぐため、
			// t.Fatal("hoge")
			t.Fatal("Wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(BitCoin(10))

		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {

		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, BitCoin(20))
		assertError(t, err, ErrInsufficientFunds)

	})

}
