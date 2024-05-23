package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Should to print hello and person's name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("want: '%s', got:'%s'", want, got)
		}
	})

	t.Run("Should to print hello world when the name is empty", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("want: '%s', got:'%s'", want, got)
		}

	})
}
