package main

import "fmt"

/*
Un profesor de programacion esta corregiendo los examenes
de sus estudiantes de la materia Programacion para poder brindarles las
correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distantas
variables. Necesita ayuda para:
1-Detectar cuales de estas variables que declaro el alumno son correctas
2-Corregir las incorrectas
*/
func main() {
	/*
		var 1nombre string Esta variable no es correcta porque las variables
		no pueden comenzar con un numero
		Lo correcto seria:
	*/
	var nombre string = "Gonzalo"
	fmt.Printf("\n%s", nombre)

	var apellido string = "Veron"
	fmt.Printf("\n%s", apellido)
	//esta variable esta bien declarada

	//Esta mal declarada
	//var int edad
	//Lo correcto:
	var edad int = 4
	fmt.Printf("\n%d",edad)

	
	//Esta mal declarada: no pudee empezar por un numero
	//1apellido := 6
	//Lo correcto:
	apellido1 := 6
	fmt.Printf("\n%d", apellido1)
	//Esta bien declarada
	var liciencia_de_conducion = true
	fmt.Printf("\n%t", liciencia_de_conducion)
	
	//Esta mal declarada
	//var estatua de la persona int
	//Lo correcto seria:
	var estatua_de_la_persona int = 4
	fmt.Printf("\n%d", estatua_de_la_persona)

	//Esta bien declarada
	cantidadDeHijos := 2
	fmt.Printf("\n%d", cantidadDeHijos)
}
