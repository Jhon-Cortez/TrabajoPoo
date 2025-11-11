using System;

namespace poocorredor
{
    //Aca vamos a ver un ejemplo de clase y uno de Herencia
    public class Nevera : Electrodomestico
    {
        private double peso;
        private double temperatura;
        private bool sirve;

        //Metodo constructor
        public Nevera(double peso, double temperatura, bool sirve)
        {
            this.peso = peso;
            this.temperatura = temperatura;
            this.sirve = sirve;
        }
        public void ahorrarEnergia(double temperatura)
        {
            if (sirve == false)
            {
                Console.WriteLine("No sirve mijo");
            }
            else if (temperatura < 0)
            {
                encender(false);
            }
            else
            {
                encender(true);
            }
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