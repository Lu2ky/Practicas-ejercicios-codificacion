package main

import "fmt"

func main() {
	var functionX func(int) int //Declaramos funcion alias con argumentos int y retorno int
	functionX = func1           //asignamos alias para func1
	fmt.Println(func2(functionX, 10))
	/*
		Aqui usaremos la functionX pero la functionX es func1
		con un alias, se lo pasaremos a la func2, la cual
		recibe una funcion como argumento y tambien un int
		el resultado de functionX  con su respectivo argumento
		se le sumara 2
	*/

}
func func1(x int) int {
	return x * 2
}
func func2(operablefunc func(int) int, y int) int {
	return operablefunc(y) + 2
}
