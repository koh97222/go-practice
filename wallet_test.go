package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(BitCoin(10))

	got := wallet.Balance()
	want := BitCoin(10)

	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
