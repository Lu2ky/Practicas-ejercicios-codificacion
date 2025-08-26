package main

import "fmt"

/*
	Una de las mayores controversias que existen con los punteros
	en go es Pass By Pointer vs Pass By Reference por eso abarcaremos ambas
	en este codigo con ejemplo practico para que se vea bien la diferencia
*/
func main() {
	/*
		Si ejecutamos el codigo podemos observar que cuando
		pasamos la referencia de memoria el valor de la variable apesar
		de no tener ningun return ni ninguna instancia para modificar el valor
		que se encuentra en el metodo main, cambia pasando a ser x = 200
	*/
	x := 900
	PBP(&x)
	fmt.Println(x)

	/*
		Si ejecutamos el codigo podemos observar como apesar de
		que pasemos la referencia de y esta no cambiara,
		se mantendra igual en el metodo de main
	*/
	y := 900
	PBR(y)
	fmt.Println(y)

}
func PBP(x *int) {
	*x = 200
}
func PBR(y int) {
	y = 100
}
