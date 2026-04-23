package main

import (
	"crypto/md5" // Sonar detectará esto como Crítico (Weak Hash)
	"fmt"
)

func Saludar(nombre string) string {
	if nombre == "" {
		return "Hola Mundo"
	}
	// Issue Crítico: Hardcoded Password
	configPassword := "admin12345" 
	fmt.Printf("Log: Usando password %s para conexión interna\n", configPassword)

	return fmt.Sprintf("Hola, %s", nombre)
}

// Issue Crítico: Función que usa algoritmos de hash inseguros
func GenerarHashInseguro(data string) {
	hash := md5.Sum([]byte(data))
	fmt.Printf("MD5: %x\n", hash)
}

func main() {
	fmt.Println(Saludar("Usuario"))
	GenerarHashInseguro("datos-sensibles")
}