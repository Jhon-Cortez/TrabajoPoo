
package poocorredor;

public abstract class Electrodomestico {
    //un ejemplo por cada tipo de dato para el ejemplo, pero igual voy a utilizar alguunos para el trabajo
    //tambien hay ejemplos de encapsulamiento al colocar algunos en public, private y protected
    protected int a = 0;
    protected byte b = 0;
    public short c = 0;
    private double d = 0;
    private float e = 0;
    private boolean encendido = false;
    protected char g = 4;
    protected long h = 0;
    protected String i = null;
    
    
    //Usamos metodos get y set para acceder y modificar a los datos
    public int getA(){
        return a;
    }
    public void setA(int a){
        this.a = a;
    }
    public int getB(){
        return b;
    }
    public void setB(byte b){
        this.b = b;
    }
    public short getC(){
        return c;
    }
    public void setC(short c){
        this.c = c;
    }
    public double getE(){
        return e;
    }
    public void setE(float e){
        this.e = e;
    }
    public boolean getEncendido(){
        return encendido;
    }
    public void setEncendido(boolean encendido){
        this.encendido = encendido;
    }
    public char getG(){
        return g;
    }
    public void setG(char g){
        this.g = g;
    }
    public long getH(){
        return h;
    }
    public void setH(long h){
        this.h = h;
    }
    
    
    
    //metodo abstracto con parametro
    public abstract void encender(boolean encendido);
    //metodo abstracto sin parametro
    public abstract void funcionar();
    //metodo abstracto con parametro
} 
