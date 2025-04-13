package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people ", func(t *testing.T) {
		got := Hello("Hamza")
		want := "Hello Hamza"

		if(got != want) {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying Hello World when emoty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if(got != want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

