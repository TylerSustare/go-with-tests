package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // is needed to tell the test suite that this method is a helper
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say Hello to people", func(t *testing.T) {
		got := Hello("You", "")
		want := "Hello, You"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
