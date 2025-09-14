package main

import (
	"fmt"
)

/*
	Las estructuras o moldes, son un tipo de estructura de datos

que nos ayuda a contener varias variables de cualquier tipo
en una sola estructura, asociadas en base al contexto
*/
type persona struct {
	Nombre string
	Edad   int
	Id     int
}

/*
Aquellas variables que su nombre empiece por mayuscula
son publicas (accesibles desde otros paquetes), aquellas
variables que empiecen por minuscula son privadas
(no accesibles desde otros paquetes unicamente desde el propio
paquete donde se instancia)
*/
func main() {
	p := persona{Nombre: "Juan", Edad: 20, Id: 912}
	fmt.Println(p.Nombre, " ", p.Edad, " ", p.Id)
	/*
		Podemos tambien cambiar las variables
		que tiene p
	*/
	p.Nombre = "Pepe"
	fmt.Println(p.Nombre, " ", p.Edad, " ", p.Id)
	/*
		Almacenar tambien las variables
		que tiene este molde
	*/
	x := p.Nombre
	fmt.Println(x)

	/*
		Inicializacion completamente
		vacia:

		p1 := new(persona)


	*/

}
