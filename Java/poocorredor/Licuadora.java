package poocorredor;

public class Licuadora extends Electrodomestico {
     private boolean sirve;
    private String marca;

    public Licuadora(boolean sirve, String marca) {
        this.sirve = sirve;
        this.marca = marca;
    }
    //Metodo para saber si la licuadora sirve o no
    public void servir(boolean sirve){
        if (sirve == false) {
            System.out.println("No sirve");
        }else{
            System.out.println("Si sirve");
        }

    //Metodos heredados de la clase electrodomestico que son abstractos

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
   
    


