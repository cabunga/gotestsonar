package main

import "testing"

func TestSaludar(t *testing.T) {
	t.Run("saludo con nombre", func(t *testing.T) {
		got := Saludar("Gemini")
		want := "Hola, Gemini"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("saludo vacío (default)", func(t *testing.T) {
		got := Saludar("")
		want := "Hola Mundo"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}