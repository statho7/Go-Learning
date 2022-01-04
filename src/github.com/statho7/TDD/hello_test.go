package main

import "testing"

func TestHello(t *testing.T) {
    // got := Hello()
    got := Hello("Chris")
    // want := "Hello, world"
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}