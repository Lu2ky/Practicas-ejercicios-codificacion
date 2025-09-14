package main

func var_declaration() {
	/*
		La declaracion de variables en go se puede hacer de dos
		formas, dinamicas y fijas.

		Dinamica:
	*/
	x := 4
	y := "ewe"
	/*
		Este tipo de declaracion funciona para que el mismo Go determine
		que tipo de variable es, en este caso seria un int debido a que
		es 4

		Fija:
	*/
	var y int = 4

	/*
		Este tipo de declaracion sirve para determinar el tipo de varible de forma
		constante, es decir no dinamica, no cambiante, se declara como int y se queda como int
	*/
}
