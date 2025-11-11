//No hay clases abstractas, esta es una simulacion de una
class Electrodomestico {
    //un ejemplo por cada tipo de dato para el ejemplo, pero igual voy a utilizar alguunos para el trabajo
    //tambien hay ejemplos de encapsulamiento al colocar algunos en public, private y protected
    constructor() {
        this.a = 0;
        this.b = 0;
        this.c = 0;
        this.d = 0;
        this.e = 0;
        this.encendido = false;
        this.g = '4';
        this.h = 0;
        this.i = null;
    }
    //los metodos get y set en JS 
    getA() { return this.a; }
    setA(value) { this.a = value; }

    getB() { return this.b; }
    setB(value) { this.b = value; }

    getC() { return this.c; }
    setC(value) { this.c = value; }

    getD() { return this.d; }
    setD(value) { this.d = value; }

    getE() { return this.e; }
    setE(value) { this.e = value; }

    getEncendido() { return this.encendido; }
    setEncendido(value) { this.encendido = value; }

    getG() { return this.g; }
    setG(value) { this.g = value; }

    getH() { return this.h; }
    setH(value) { this.h = value; }

    getI() { return this.i; }
    setI(value) { this.i = value; }

    //metodos abstratos simulados
    encender(encendido) {
       console.log("El electrodomestico se ha encendido: " + encendido);
    }

    funcionar() {
        console.log("El electrodomestico esta funcionando");
    }
}
// Se utiliza para exponer la clase en otros archivos
module.exports = Electrodomestico;
