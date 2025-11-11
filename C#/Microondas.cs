using System;

namespace poocorredor
{
    public class Microondas : Electrodomestico
    {
        private bool sirve;
        private string marca;

        public Microondas(bool sirve, string marca)
        {
            this.sirve = sirve;
            this.marca = marca;
        }
        //Metodos heredados de la clase electrodomestico que son abstractos
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
