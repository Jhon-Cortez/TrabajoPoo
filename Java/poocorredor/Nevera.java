/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package poocorredor;

/**
 *
 * @author LENOVO
 */
//Aca vamos a ver un ejemplo de clase y uno de Herencia
public class Nevera extends Electrodomestico{

    private double peso;
    private double temperatura;
    private boolean sirve;
    
    public Nevera(double peso, double temperatura, boolean sirve){
        this.peso = peso;
        this.temperatura = temperatura;
        this.sirve = sirve;
    }
    
    public void ahorrarEnergia(double temperatura){
        if (sirve == false) {
            System.out.println("No sirve mijo");
        }else if (temperatura<0) {
            encender(false);
        }else{
            encender(true);
        }
    }
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
