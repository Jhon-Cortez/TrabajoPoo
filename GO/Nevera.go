package main

import "fmt"
type Nevera struct {
	Electrodomestico
	Capacidad   float64
	Temperatura float64
	Sirve       bool
}
//Se crea el metodo constructor de nevera
func NewNevera(capacidad float64, temperatura float64, sirve bool) *Nevera {
	return &Nevera{
		Capacidad:   capacidad,
		Temperatura: temperatura,
		Sirve:       sirve,
	}
}
// Implementamos los métodos abstractos
func (n *Nevera) Encender(encendido bool) {
	n.encendido = encendido
	if encendido {
		fmt.Println("Encendido")
	} else {
		fmt.Println("Apagado")
	}
}

func (n Nevera) Funcionar() {
	if n.encendido {
		fmt.Println("La nevera esta funcionando")
	} else {
		fmt.Println("La nevera no esta funcionando")
	}
}