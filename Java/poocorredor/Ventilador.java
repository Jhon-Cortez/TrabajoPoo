package poocorredor;
//Se extiende de la clase electrodomestico
public class Ventilador extends Electrodomestico {

    private double velocidad;
    private boolean sirve;
    //Constructor del ventilador
    public Ventilador(double velocidad, boolean sirve){
        this.velocidad = velocidad;
        this.sirve = sirve;
    }
    //Metodo para saber si el ventilador sirve o no
    public void servir(double velocidad){
        if (sirve == false) {
            System.out.println("No sirve");
        }else if (velocidad>3) {
            encender(false);
        }else{
            encender(true);
        }
    }
    //Metodo para saber si el ventilador esta tan rapido que puede volar
    public void volar(double velocidad){
        if (velocidad > 80) {
            System.out.println("El ventilador empieza a volar");
        }
        
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
