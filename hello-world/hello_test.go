package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people ", func(t *testing.T) {
		got := Hello("Hamza")
		want := "Hello Hamza"

		assertCorrectMessage(t, got, want);
	})

	t.Run("saying Hello World when empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		assertCorrectMessage(t, got, want);

	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
