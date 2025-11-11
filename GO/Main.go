//El paquete main debe estar en el resto de clases para que no de errores
package main
import "fmt"

func main() {
    fmt.Println("Holasaa")
    e := NewEstufa(12.3, 200.0, true)//Aca esta el metodo abstracto
	e.Sirve = false
	fmt.Println(e)
    e.Encender(true)
}