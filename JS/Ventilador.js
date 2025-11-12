const Electrodomestico = require("./Electrodomestico");

//Se extiende de la clase electrodomestico
class Ventilador extends Electrodomestico {
    //Constructor del ventilador
    constructor(velocidad, sirve) {
        super();
        this.velocidad = velocidad;
        this.sirve = sirve;
    }
    //Metodo para saber si el ventilador sirve o no
    servir(velocidad) {
        if (this.sirve === false) {
            console.log("No sirve");
        } else if (velocidad > 3) {
            this.encender(false);
        } else {
            this.encender(true);
        }
    }

    //Metodo para saber si el ventilador esta tan rapido que puede volar
    volar(velocidad) {
        if (velocidad > 80) {
            console.log("El ventilador empieza a volar");
        }
    }

    //Metodos heredados de la clase electrodomestico que son abstractos
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

module.exports = Ventilador;