package main

import (
	"errors"
	"fmt"
)

// BitCoin BitCoin
type BitCoin int

// Wallet 財布
type Wallet struct {
	balance BitCoin
}

// Stringer インターフェース
type Stringer interface {
	String() string
}

// ErrInsufficientFunds error message when withdraw over balance.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Deposit デポジット
func (w *Wallet) Deposit(amount BitCoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance 残高
func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw Withdraw
func (w *Wallet) Withdraw(amount BitCoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
