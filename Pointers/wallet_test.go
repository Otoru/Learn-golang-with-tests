package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleWallet() {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	fmt.Println(wallet.Balance())
	// Output: 10 BTC
}

func BenchmarkWallet(b *testing.B) {
	wallet := Wallet{}
	bitcoins := Bitcoin(rand.Intn(9))

	for i := 0; i < b.N; i++ {
		wallet.Deposit(bitcoins)
	}

	fmt.Println(wallet.Balance())
}

func TestWallet(t *testing.T) {

	t.Run("Test deposit function", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Test withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Test withdraw without sufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()
	assertValue(t, result, expected)
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but didn't want one.")
	}
}

func assertError(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("Didn't get an error but wanted one.")
	}
	assertValue(t, result, expected)
}

func assertValue(t *testing.T, result, expected interface{}) {
	t.Helper()
	if result != expected {
		t.Errorf("got '%q' and want '%q'.", result, expected)
	}
}
