package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hamza")
	want := "Hello Hamza"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

