package main

import "testing"

func TestSearch(t *testing.T){
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"


	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}
	assertMessage(t, got, want)
}