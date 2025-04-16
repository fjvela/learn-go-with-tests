package pointers

type Wallet struct{}

func (w *Wallet) Deposit(amount float64) {
}

func (w *Wallet) Balance() float64 {
	return 0
}
