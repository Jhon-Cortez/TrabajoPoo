using System;

namespace poocorredor
{
    //clase abstracta como en java
    public abstract class Electrodomestico
    {
        //un ejemplo por cada tipo de dato para el ejemplo, pero igual voy a utilizar alguunos para el trabajo
        //tambien hay ejemplos de encapsulamiento al colocar algunos en public, private y protected
        protected int a = 0;
        protected byte b = 0;
        public short c = 0;
        private double d = 0;
        private float e = 0;
        private bool encendido = false;
        protected char g = '4';
        protected long h = 0;
        protected string i = null;

        //constructor vacio
        public Electrodomestico() { }

        //getters y setters 
        public int GetA() { return a; }
        public void SetA(int value) { a = value; }

        public byte GetB() { return b; }
        public void SetB(byte value) { b = value; }

        public short GetC() { return c; }
        public void SetC(short value) { c = value; }

        public double GetD() { return d; }
        public void SetD(double value) { d = value; }

        public float GetE() { return e; }
        public void SetE(float value) { e = value; }

        public bool GetEncendido() { return encendido; }
        public void SetEncendido(bool value) { encendido = value; }

        public char GetG() { return g; }
        public void SetG(char value) { g = value; }

        public long GetH() { return h; }
        public void SetH(long value) { h = value; }

        public string GetI() { return i; }
        public void SetI(string value) { i = value; }

        //metodos abstractos
        public abstract void encender(bool encendido);

        public abstract void funcionar();
    }
}