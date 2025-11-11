package main
type Electrodomestico struct {
	//protected y privado es igual
	a int       // "protected" en minúscula
    b byte      // igual
    C int16     // "public" en mayúscula
    d float64   // "private" en minúscula
    e float32
    encendido bool
    g rune      // rune es char en Go 
    h int64     // equivalente a long
    i string    // equivalente a String
	}
	//Metodos get y sett
	func (e Electrodomestico) A() int {
	return e.a
	}
	func (e *Electrodomestico) SetA(a int) {
		e.a = a
	}

	func (e Electrodomestico) B() byte {
		return e.b
	}
	func (e *Electrodomestico) SetB(b byte) {
		e.b = b
	}

	func (e Electrodomestico) C_() int16 { //Se usa C_ por que C ya esta en uso por que se coloco publico
		return e.C
	}
	func (e *Electrodomestico) SetC(c int16) {
		e.C = c
	}

	func (e Electrodomestico) E() float32 {
		return e.e
	}
	func (e *Electrodomestico) SetE(v float32) { // Aca igualmente no se usa e por que ya esta usado entonces se usa una variable que no se halla definido antes
		e.e = v
	}

	func (e Electrodomestico) Encendido() bool {
		return e.encendido
	}
	func (e *Electrodomestico) SetEncendido(encendido bool) {
		e.encendido = encendido
	}

	func (e Electrodomestico) G() rune {
		return e.g
	}
	func (e *Electrodomestico) SetG(g rune) {
		e.g = g
	}

	func (e Electrodomestico) H() int64 {
		return e.h
	}
	func (e *Electrodomestico) SetH(h int64) {
		e.h = h
	}
