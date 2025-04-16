package pointers

import (
	"fmt"
	"testing"
)

func TestWaller(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(11.0)

	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
