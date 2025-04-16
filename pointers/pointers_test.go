package pointers

import (
	"testing"
)

func TestWaller(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, msgExpected string) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err.Error() != msgExpected {
			t.Errorf("got %q want %q", err.Error(), msgExpected)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10.0)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		error := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10.0)

		assertNoError(t, error)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw no enough funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw 100.00 BTC, insufficient funds: 20.00 BTC")
		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
