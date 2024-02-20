package pointerrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.00000523456789)) // max decimals able
		want := Bitcoin(10.00000523456789)

		assertBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.00000523456789)}
		err := wallet.Withdraw(Bitcoin(3.5))
		want := Bitcoin(6.50000523456789)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(35))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but did not want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}
