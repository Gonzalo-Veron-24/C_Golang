/*
1. Crea una aplicación donde tengas como variable tu nombre y dirección.
2. Imprime en consola el valor de cada una de las variables.
*/

package main

import "fmt"

func main() {
	var (
		nombre = "Veron Gonzalo"
		dirc   = "Mexico 367"
	)
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Direccion:  %s\n", dirc)
}
