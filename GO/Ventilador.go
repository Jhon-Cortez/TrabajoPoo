package main
import "fmt"
type Ventilador struct {
	Electrodomestico
	Velocidad int
	Sirve     bool
}
//Se crea el metodo constructor de ventilador
func NewVentilador(velocidad int, sirve bool) *Ventilador {
	return &Ventilador{
		Velocidad: velocidad,
		Sirve:     sirve,
	}
}
// Implementamos los métodos abstractos
func (v *Ventilador) Encender(encendido bool) {
	v.encendido = encendido
	if encendido {
		fmt.Println("💨 Ventilador encendido")
	} else {
		fmt.Println("Ventilador apagado")
	}
}

func (v Ventilador) Funcionar() {
	if v.encendido {
		fmt.Println("El ventilador esta ventilando el aire")
	} else {
		fmt.Println("El ventilador no funciona")
	}
}