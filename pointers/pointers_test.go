package pointers

import (
	"fmt"
	"testing"
)

func TestWaller(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.0

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
