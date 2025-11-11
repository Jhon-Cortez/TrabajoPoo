using System;

namespace poocorredor
{
    public class Licuadora : Electrodomestico
    {
        //atributos de licuadora
        private string marca;
        private int velocidad;

        //constructor de licuadora
        public Licuadora(string marca, int velocidad)
        {
            this.marca = marca;
            this.velocidad = velocidad;
        }

        public override void encender(bool encendido)
        {
            if (encendido == true)
            {
                Console.WriteLine("Se enciende");
            }
            else
            {
                Console.WriteLine("se apaga");
            }
        }

        public override void funcionar()
        {
            Console.WriteLine("Funciono");
        }
    }
}