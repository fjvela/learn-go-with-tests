package pointers

import "fmt"

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// https://go.dev/ref/spec#Method_values
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return fmt.Errorf("cannot withdraw %s, insufficient funds: %s", amount, w.balance)
	}

	w.balance -= amount
	return nil
}
