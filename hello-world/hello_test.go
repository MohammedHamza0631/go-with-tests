package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people ", func(t *testing.T) {
		got := Hello("Hamza", "")
		want := "Hello Hamza"

		assertCorrectMessage(t, got, want);
	})

	t.Run("saying Hello World when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want);

	})

	t.Run("saying Hello in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello in French", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour Chris"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper();
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
