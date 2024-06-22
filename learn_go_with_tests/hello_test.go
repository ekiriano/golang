package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Amine", "")
		want := "Hello, Amine"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Amine", "French")
		want := "Bonjour, Amine"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Writing a test is just like writing a function, with a few rules

// It needs to be in a file with a name like xxx_test.go

// The test function must start with the word Test

// The test function takes one argument only t *testing.T

// To use the *testing.T type, you need to import "testing", like we did with fmt in the other file
