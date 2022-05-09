/*
Una empresa de meteorología quiere tener una aplicación donde pueda tener la 
temperatura y humedad y presión atmosférica de distintos lugares. 
1. Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, 
humedad y presión de donde te encuentres. 
2. Imprime los valores de las variables en consola. 
3. ¿Qué tipo de dato le asignarías a las variables?

*/

package main
import "fmt"

func main(){
// Consigna 1
	var(
		Temperatura float32 = 24.5
		Humedad float32 = 9.5
		Presion float32 = 1014
	)
//Consigna 2
	fmt.Printf("Temperatura: %.2f\n",Temperatura)
	fmt.Printf("Humedad: %.2f\n",Humedad)
	fmt.Printf("Presion: %.2f\n",Presion)
}