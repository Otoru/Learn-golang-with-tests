package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleHello() {
	message := Hello("Vitor", "english")
	fmt.Println(message)
	// Output: Hello Vitor
}

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%q' and want '%q'.", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Vitor", "english")
		const want string = "Hello Vitor"
		assertMessage(t, got, want)
	})

	t.Run("Say 'Hello World' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("", "English")
		const want string = "Hello World"
		assertMessage(t, got, want)
	})

	t.Run("Saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Vitor", "Spanish")
		const want string = "Hola Vitor"
		assertMessage(t, got, want)
	})

	t.Run("Saying hello to people in French", func(t *testing.T) {
		got := Hello("Vitor", "French")
		const want string = "Bonjour Vitor"
		assertMessage(t, got, want)
	})

}

func BenchmarkHello(b *testing.B) {
	options := []string{"English", "Spanish", "French"}
	len := len(options)
	for i := 0; i < b.N; i++ {
		choice := rand.Intn(len)
		Hello("Vitor", options[choice])
	}
}
