
package poocorredor;

public class Microondas extends Electrodomestico {
    private boolean sirve;
    private String marca;

    public Microondas(boolean sirve, String marca) {
        this.sirve = sirve;
        this.marca = marca;
    }

    
    //Metodos heredados de la clase electrodomestico que son abstractos
    @Override
    public void encender(boolean encendido) {
        if (encendido==true) {
            System.out.println("Se enciende");
        }else{
            System.out.println("se apaga");
        }
    }
    @Override
    public void funcionar() {
        System.out.println("Funciono");
        
    }
    
}
