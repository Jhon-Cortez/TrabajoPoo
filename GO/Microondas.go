package main

import "fmt"
type Microondas struct {
	Electrodomestico
	Potencia int
	Sirve    bool
}
//Se crea el metodo constructor de microondas
func NewMicroondas(potencia int, sirve bool) *Microondas {
	return &Microondas{
		Potencia: potencia,
		Sirve:    sirve,
	}
}
// Implementamos los métodos abstractos
func (m *Microondas) Encender(encendido bool) {
	m.encendido = encendido
	if encendido {
		fmt.Println("encendido")
	} else {
		fmt.Println("Microondas apagado")
	}
}

func (m Microondas) Funcionar() {
	if m.encendido {
		fmt.Println("El microondas está calentando la comida")
	} else {
		fmt.Println("El microondas no funciona")
	}
}
