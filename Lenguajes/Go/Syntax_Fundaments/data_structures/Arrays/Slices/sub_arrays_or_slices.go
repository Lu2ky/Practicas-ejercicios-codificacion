package main

import "fmt"

func main() {
	/*
		Un slice, es el corte o una vista de un array, es decir
		el partido de un array, todos los cambios que se hagan en el
		slice se almacenan en el array original, es una estructura de
		dato muy util para algoritmos de tipo "divide y venceras", donde buscamos
		siempre dividir el problema

		Metafora:
		Digamos que tienes un tomate, ese tomate lo puedes cortar en dos, esos dos los puedes cortar en
		otros dos pedazos, pasa exactamente lo mismo con los arrays donde un array es un tomate

		El slice se compone de dos cosas, el puntero y la longitud, el punto que determina el punto donde
		se desea cortar y la longitud hasta donde

		Ejemplo:

		[1][2][3][4][5][6][7] <------ Valores
		 |  |  |  |  |  |  |
		 0  1  2  3  4  5  6 <------- Indeces


		Puntero: 2, longitud: 4
		Slice resultante:
		[3][4][5][6]

		Declaracion de un slice

		var slice []T
		slice := make([]T, 0)

		Desglosado>> Inicializamos un array, luego ese array sera igual a el retorno del metodo "make"
					 el cual recibe un tipo de elemento []T,numero incial de numeros y capacidad, este ultimo
					 es opcional si no se pone sera igual al numero inicial de numeros, un slice siempre tendra
					 estas 3 cosas, puntero, longitud y capacidad

		Aplicado:
	*/
	var x [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	slice1 := x[0:3]
	slice2 := x[4:7]
	fmt.Println(slice1)
	fmt.Println(slice2)

	/*
		Partimos el array en dos pedazos, el pedazos de el indice 0 a 2 y el de  4 a 6
		los slices tienen tambien la propiedad de capacidad que mencionamos anteriormente,
		la capacidad es la longitud hipotetica que puede llegar a tener un slice donde se toma
		desde el puntero hasta el len del array cortado

		Formula:
				cap(slice)=len(array)−puntero

		Una forma de verlo: la capacidad son los elementos del array que no
		se ven directamente en el slice, pero que sí se afectan si modificamos o expandimos
		el slice con append, mientras no se supere esa capacidad.
	*/
	fmt.Println("Capacidad slice1: ", cap(slice1))
	fmt.Println("Longitud slice1: ", len(slice1))
	/*
		El len es la cantidad de datos que contiene en los espacios donde
		no es nil(null en otros lenguajes)


		Funcion append esta funcion nos ayuda a agregar elementos a un slice
		si el slice llega a su capacidad maxima este buscara aumentarla

		IMPORTANTE:

		Los cambios que se hagan con append se reflejaran en el array padre mientras
		y solo mientras no se supere la capacidad del slice
	*/
	slice2 = append(slice2, 200)
	fmt.Println(slice2)

}
