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
	var y [5]int = [5]{1,2,3,4,5}
	/*
	Los arrays pueden variar en dimensiones 2,3,4,5,6,etc.
	*/
	var y1 [][]int
	var y2 [][][]int
	// Asi sucesivamente...

	/* 

	*/

}
