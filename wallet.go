package main

import "fmt"

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
