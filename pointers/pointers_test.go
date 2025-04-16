package pointers

import (
	"testing"
)

func TestWaller(t *testing.T) {

	assertBalance := func(t testing.TB, got Bitcoin, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10.0)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10.0)

		assertBalance(t, got, want)
	})
}
