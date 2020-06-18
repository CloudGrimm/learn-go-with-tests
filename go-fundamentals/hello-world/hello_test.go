package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Wilbrod")
	want := "Hello, Wilbrod"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
