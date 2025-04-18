package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Javi")

	got := buffer.String()
	want := "Hello, Javi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
