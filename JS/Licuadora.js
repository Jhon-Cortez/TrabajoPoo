//Aca es para poder llamar la clase padre y para eso fue que la exposimos
const Electrodomestico = require("./Electrodomestico");

//Se extiende igual que en Java
class Licuadora extends Electrodomestico {
    //atributos de esta clase
    constructor(marca, velocidad) {
        //El super es para llamar al constructor de la clase padre
        super();
        this.marca = marca;
        this.velocidad = velocidad;
    }
    getVelocidad() { return this.velocidad; }
    setVelocidad(value) { this.velocidad = value; }
    //Metodos "abstractos" 
    encender(encendido) {
        if (encendido === true) {
            console.log("Se enciende");
        } else {
            console.log("se apaga");
        }
    }

    funcionar() {
        console.log("Funciono");
    }
}

module.exports = Licuadora;
