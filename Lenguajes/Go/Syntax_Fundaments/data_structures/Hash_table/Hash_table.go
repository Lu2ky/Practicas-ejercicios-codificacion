package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		Las tablas hash son pares de valores y llaves relacionados entre si
		para ingresar a un valor se necesita un llave, la hash table tiene n
		espacios de almacenamiento donde la posicion se determina con
		la funcion hash de la llave y posteriormente se saca el modulo con los
		n espacios, esto para que se mantenga en el rango n

		posicion = hashkey%n


		Declaracion:
		var HashTable map[K]V

		K y V hacen referencia a tipos de variables, vea Var_declaration/var_types

	*/
	var ExampleMap map[string]int     //Llave de tipo string y valor de tipo int, inicializado en nil
	ExampleMap = make(map[string]int) //Crea la estructura interna del mapa y devulve la referencia

	/*
		Para inicializar un mapa ya con valores se hace de esta forma
	*/
	Example2Map := map[string]int{
		"Hola": 901}
	/*
		Para sacar valores de este mapa debemos utilizar la llave
		de la siguiente manera
	*/
	fmt.Println(Example2Map["Hola"])

	ExampleMap["Juan"] = 812
	fmt.Println(ExampleMap["Juan"])

	/*
		Para los mapas tambien podemos tener
		un doble output, el primero sera el valor
		y el segundo un booleano si se
		encuentra o no en el mapa
	*/
	_, b := ExampleMap["Javier"]
	fmt.Println("Encontrado?: ", b)

	/*
		Los mapas tambien tienen longitud,
		se utiliza len(NombreMapa)
	*/
	fmt.Println("Tamaño mapa: ", len(ExampleMap))
	/*
		Agregaremos algunos valores al ExampleMap
		para ver como se recorre y lo recorreremos posteriormente
	*/
	ExampleMap["Ana"] = rand.Intn(101)
	ExampleMap["Luis"] = rand.Intn(101)
	ExampleMap["Pedro"] = rand.Intn(101)
	ExampleMap["Maria"] = rand.Intn(101)
	ExampleMap["Carlos"] = rand.Intn(101)
	ExampleMap["Elena"] = rand.Intn(101)
	ExampleMap["Jorge"] = rand.Intn(101)
	ExampleMap["Sofia"] = rand.Intn(101)
	ExampleMap["Andres"] = rand.Intn(101)
	ExampleMap["Valeria"] = rand.Intn(101)

	for key, val := range ExampleMap {
		fmt.Println("Llave: ", key, " Valor: ", val)
	}
}
