package main

import "fmt"

func main() {
	/*
		Los punteros de memoria o memory pointers sirven para que una variable
		apunte directamente al lugar donde se almacena un dato en la memoria
		la sintaxis para declarar un puntero es la siguiente(INTEGER):
	*/
	var pointer *int
	//var pointer *string
	/*
		La variable pointer es una variable de tipo pointer a int, esto quiere decir
		que apuntara a valores int, incializaremos una variable x en 1 para hacer un ejemplo
	*/
	var x int = 1
	/*
		Apuntamos hacia la variable x con el pointer declarado en la linea 11
	*/
	pointer = &x
	fmt.Println(pointer)
	/*
		Imprime la direccion de memoria debido a que nos falta desreferenciar este puntero, inicializaremos
		una variable y y posteriormente desreferenciaremos el puntero, esto para mostrar la variable a la que corresponde
	*/
	y := *pointer
	fmt.Println(y)

}
