using System;

namespace poocorredor
{
    public class Estufa : Electrodomestico
    {
        //atributos de la clase Estufa
        private double peso;
        private double temperatura;
        private bool sirve;

        //Metodo constructor en c#
        public Estufa(double peso, double temperatura, bool sirve)
        {
            peso = peso;
            temperatura = temperatura;
            sirve = sirve;
        }


        //Metodo para que no halla sobrecalentamiento
        public void prevencion(double temperatura)
        {
            if (sirve == false)
            {
                System.Console.WriteLine(); ("No sirve mijo");
            }
            else if (temperatura > 200)
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
