const Electrodomestico = require("./Electrodomestico");

class Estufa extends Electrodomestico {
    //atributos de la clase
    constructor(peso, temporada, gas) {
        super();
        this.peso = peso;
        this.sirve = sirve;
    }
    //los metodos abstractos
    prevencion(temperatura){
        if (sirve == false) {
                console.log("No sirve");
        }else if (temperatura>200) {
                encender(false);
        }else{
                encender(true);
        }
    }
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

module.exports = Estufa;
