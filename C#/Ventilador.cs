using System;

namespace poocorredor
{
    //Se extiende de la clase electrodomestico
    public class Ventilador : Electrodomestico
    {
        private double velocidad;
        private bool sirve;

        //Constructor del ventilador
        public Ventilador(double velocidad, bool sirve)
        {
            this.velocidad = velocidad;
            this.sirve = sirve;
        }

        //Metodo para saber si el ventilador sirve o no
        public void servir(double velocidad)
        {
            if (sirve == false)
            {
                Console.WriteLine("No sirve");
            }
            else if (velocidad > 3)
            {
                encender(false);
            }
            else
            {
                encender(true);
            }
        }

        //Metodo para saber si el ventilador esta tan rapido que puede volar
        public void volar(double velocidad)
        {
            if (velocidad > 80)
            {
                Console.WriteLine("El ventilador empieza a volar");
            }
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
