package main

import "fmt"

type Estufa struct {
	Electrodomestico // Aca se hace la "herencia"(Es una composicion)
	Peso             float64
	Temperatura      float64
	Sirve            bool
}

// Constructor (Go no tiene constructores automáticos)
func NewEstufa(peso float64, temperatura float64, sirve bool) *Estufa {
	return &Estufa{
		Peso:        peso,
		Temperatura: temperatura,
		Sirve:       sirve,
	}
}
func (e *Estufa) Prevencion(temp float64) {
	if !e.Sirve {
		fmt.Println("No sirve mijo")
	} else if temp > 200 {
		e.Encender(false)
	} else {
		e.Encender(true)
	}
}

// Sobrescribimos Encender (polimorfismo)
func (e *Estufa) Encender(encendido bool) {
	e.encendido = encendido
	if encendido {
		fmt.Println("Se enciende")
	} else {
		fmt.Println("Se apaga")
	}
}

// Implementa el otro método abstracto
func (e Estufa) Funcionar() {
	fmt.Println("Funciono")
}
