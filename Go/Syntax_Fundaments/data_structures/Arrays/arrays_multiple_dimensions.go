package main

import "fmt"

func main() {
	/*
		Los arrays son un tipo de structura de dato que estan relacionados en una lista que
		posee un indice y un valor, para inicializar un array se hace de la siguiete manera:
	*/
	var x [5]int
	/*
		Este es un array que tiene 5 espacios, siendo 0,1,2,3,4 en sus indices, para introducir
		datos a este array tenemos que escribir el nombre del array y posteriormente el indice
		y el valor, ingresaremos un int 4 en el indice 1
		[][4][][][]
		Aqui se muestra como se estructura el array despues de el ingreso del numero 4
	*/
	x[1] = 4
	fmt.Println(x[1])
	/*
		Arrays inicializados y rellenados, estos arrays son aquellos que son rellenados en el mismo momentos
		que son inicializados, la sintaxis es algo asi:
	*/
	var y [5]int = [5]int{1, 2, 3, 4, 5}
	/*
		Los arrays pueden variar en dimensiones 2,3,4,5,6,etc.
	*/
	//var y1 [1][1]int = [1][1]int {1,1}
	//var y2 [][][]int = [1][1][1]int {1, 1, 1}
	// Asi sucesivamente...

	/*
		Ejemplo de array inicializado y rellenado de dos dimensiones
	*/
	var y3 [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	/*
		¿Como recorremos un array?
		Para recorrer un array tenemos en cuenta la dimension que tiene
		dependiendo de esto sera la cantidad de for que tendremos que hacer
		Ejemplo para y (array de una dimension ya declarado anteriormente)
	*/
	//For tipo 1
	for _, z := range y {
		fmt.Print(z)
	}
	fmt.Println()
	//For tipo 2
	for i, z := range y {
		fmt.Print(z)
		if i == len(y)-1 {
			fmt.Println()
			fmt.Print(i) //i, iterator
		}

	}
	fmt.Println()
	//For tipo 3
	for i := 0; i < len(y); i++ {
		fmt.Print(y[i])
	}
	/*
		La diferencia entre el for tipo 1 y tipo 2 es que el tipo 1
		no almacena una variable contador, mientras que el segundo tipo si lo hace


		Para recorrer un for de dos dimensiones tendremos que hacer la cantidad de for
		de las dimensiones que tiene, osea dos for anidados de esta manera
	*/
	fmt.Println()
	for _, fila := range y3 {
		for _, valor := range fila {
			fmt.Print(valor, " ")
		}
		fmt.Println()
	}

	/*
		Se puede tambien usar el almacenamiento declarando la variable i como primer iterador y j
		como segundo iterador si se necesita
	*/
	fmt.Println("Segunda forma: ")
	for i := 0; i < len(y3); i++ {
		for j := 0; j < len(y3[i]); j++ {
			fmt.Print(y3[i][j], " ")
		}
		fmt.Println()
	}

}
