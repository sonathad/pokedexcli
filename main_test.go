package main

import "testing"

func TestMain(t *testing.T) {
	got := Greeter()
	want := "Welcome to Pokedex CLI!"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
