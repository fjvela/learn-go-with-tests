package pointers

import "testing"

func TestWaller(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.0

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
