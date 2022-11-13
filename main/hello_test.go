package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, Chris' when Spanish is supplied", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Konnichiwa, Chris' when Japanese is supplied", func(t *testing.T) {
		got := Hello("Chris", "Japanese")
		want := "Konnichiwa, Chris"
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
