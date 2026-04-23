package main

import "fmt"

// Saludar devuelve un mensaje de bienvenida. 
// Separamos la lógica para poder testearla.
func Saludar(nombre string) string {
	if nombre == "" {
		return "Hola Mundo"
	}
	return fmt.Sprintf("Hola, %s", nombre)
}

func main() {
	fmt.Println(Saludar("Usuario"))
}