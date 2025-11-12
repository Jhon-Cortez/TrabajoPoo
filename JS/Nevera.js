const Electrodomestico = require("./Electrodomestico");

//Aca vamos a ver un ejemplo de clase y uno de Herencia
class Nevera extends Electrodomestico {
    constructor(peso, temperatura, sirve) {
        super();
        this.peso = peso;
        this.temperatura = temperatura;
        this.sirve = sirve;
    }

    //Metodo para ahorrar energia :D
    ahorrarEnergia(temperatura) {
        if (this.sirve === false) {
            console.log("No sirve");
        } else if (temperatura < 0) {
            this.encender(false);
        } else {
            this.encender(true);
        }
    }
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

module.exports = Nevera;
