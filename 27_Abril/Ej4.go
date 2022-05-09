/*
Un estudiante de programación intentó realizar declaraciones de variables con sus respectivos 
tipos en Go pero tuvo varias dudas mientras lo hacía. A partir de esto, nos brindó su código y 
pidió la ayuda de un desarrollador experimentado que pueda: 
● Verificar su código y realizar las correcciones necesarias. 
*/
package main
import "fmt"

func main(){
	var apellido string = "Gomez"
	fmt.Printf("\n%s",apellido)
	//var edad int = "35" esta definiendo un string como entero
	var edad int = 22 //Esta manera seria la correcta
	fmt.Printf("\n%d",edad)
	boolean := "false"; //con esto esta definiendo un string no un booleano
	fmt.Printf("\n%s",boolean)
	//var sueldo string = 45857.90 esta definiendo un double como string
	var sueldo float32 = 45857.90
	fmt.Printf("\n%g",sueldo)
	var nombre string = "Julian"
	fmt.Printf("\n%s",nombre)	
}
