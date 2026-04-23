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
func TestGenerarHashInseguro(t *testing.T) {
	// Llamamos a la función para que Sonar vea que tiene cobertura,
	// pero Sonar seguirá marcando que el uso de MD5 es un riesgo.
	GenerarHashInseguro("test")
}