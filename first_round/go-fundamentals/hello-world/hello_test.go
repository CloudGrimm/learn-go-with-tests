package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	//test in english
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Wilbrod", "English")
		want := "Hello, Wilbrod"
		assertCorrectMessage(t, got, want)
	})
//test for an empty string
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	//test for spanish
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Wilbrod", "Spanish")
		want := "Hola, Wilbrod"
		assertCorrectMessage(t, got, want)
	})
	//test for French
	t.Run("in French", func(t *testing.T) {
		got := Hello("Wilbrod", "French")
		want := "Bonjour, Wilbrod"
		assertCorrectMessage(t, got, want)
	})

}
