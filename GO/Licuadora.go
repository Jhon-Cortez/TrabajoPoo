package main
import "fmt"
type Licuadora struct {
	Electrodomestico
	Velocidad int
	Sirve     bool
}
//Se crea el motodo contructor de licaudora
func NewLicuadora(velocidad int, sirve bool) *Licuadora {
	return &Licuadora{
		Velocidad: velocidad,
		Sirve:     sirve,
	}
}
// Implementamos los métodos abstractos
//Metodo para decir que se encendio
func (l *Licuadora) Encender(encendido bool) {
	l.encendido = encendido
	if encendido {
		fmt.Println("Encendida")
	} else {
		fmt.Println("Apagafo")
	}
}
//metodo para decir que esta funcionando
func (l Licuadora) Funcionar() {
	if l.encendido {
		fmt.Println("La licuadora esta licuando")
	} else {
		fmt.Println("La licuadora no esta licuando")
	}
}
