package main

import "testing"

func TestHello(t *testing.T) {
	checksCorrectMessage := func(t *testing.T, want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	}
	t.Run("Should to print hello and person's name", func(t *testing.T) {
		got := Hello("PTBR", "Chris")
		want := "Olá, Chris"
		checksCorrectMessage(t, want, got)
	})

	t.Run("Should to print hello world when the name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		checksCorrectMessage(t, want, got)
	})
}

func TestLang(t *testing.T) {
	checksCorrectMessage := func(t *testing.T, want, got string) {
		t.Helper()
		if want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	}

	t.Run("Should to print Hello in PTBR", func(t *testing.T) {
		got := (Hello("PTBR", " "))
		want := "Olá, World"
		checksCorrectMessage(t, want, got)

	})

	t.Run("Should to print Hello in ESP", func(t *testing.T) {
		got := (Hello("ESP", " "))
		want := "Hola, World"
		checksCorrectMessage(t, want, got)

	})

	t.Run("Should to print Hello is Empty", func(t *testing.T) {
		got := (Hello(" ", " "))
		want := "Hello, World"
		checksCorrectMessage(t, want, got)

	})
}
