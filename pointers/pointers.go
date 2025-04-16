package pointers

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() float64 {
	// https://go.dev/ref/spec#Method_values
	return w.balance
}
