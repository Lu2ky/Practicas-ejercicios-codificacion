package main

import "fmt"

/*
	Las funciones se pueden tratar como cualquier otro tipo
	de variable, como un int, string, etc.
	Las variables pueden ser declaradas como un tipo de funcion,
	En la programacion dinamica las variables se pueden crear
	de forma dinamica, las funciones como variables se pueden
	pasar como argumentos y pueden ser retornadas como valores,
	tambien pueden almacenarse en las strucs
*/
var Var1 func(int) int

/*
	Primero declaramos la variable y el nombre, luego decimos el tipo,
	decimos que es una funcion que recibe int y que devuelve un int
*/
func incFn(x int) int {
	return x + 1
}

/*
	Definimos la funcion a la que le pondremos el alias
*/

func main() {
	Var1 = incFn       //asignamos alias
	fmt.Print(Var1(1)) //Var1 es lo mismo que incFn
}
