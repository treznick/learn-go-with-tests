package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	greet := NewGreet(Options(Output(&buffer)))

	greet.Hello("Chris")
	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
