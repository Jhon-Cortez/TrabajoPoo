package main
import "fmt"
//En go no hay metodos abstractos, toca con una interfaz
type IElectrodomestico interface {
	Encender(encendido bool)
	Funcionar()
}
type Lavadora struct {
	Electrodomestico // Herencia pero en go (Es como una implementacion en lenguajes como java)
}
// Implementamos los métodos abstractos
func (l *Lavadora) Encender(encendido bool) {
	l.encendido = encendido
	if encendido {
		fmt.Println("Lavadora encendida")
	} else {
		fmt.Println("Lavadora apagada")
	}
}
func (l Lavadora) Funcionar() {
	if l.encendido {
		fmt.Println("La lavadora funciona")
	} else {
		fmt.Println("No Funciona")
	}
}