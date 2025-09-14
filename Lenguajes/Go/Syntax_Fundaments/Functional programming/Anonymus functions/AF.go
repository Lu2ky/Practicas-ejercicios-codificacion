package main

import (
	"fmt"
)

/*
	Las funciones no necesitan definicion
	previa para ser utilizadas, se pueden instanciar
	en los mismos argumentos sin ningun tipo de nombre
*/

func main() {
	v := func2(func(x int) int {
		return x * 2
	}, 10)
	fmt.Println(v)
}
func func2(operablefunc func(int) int, y int) int {
	return operablefunc(y) + 2
}
