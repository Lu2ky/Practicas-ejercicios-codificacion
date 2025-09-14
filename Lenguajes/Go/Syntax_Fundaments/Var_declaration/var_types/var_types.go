package main

func main() {
	/*
		En Golang existen los tipos de variables mas conocidos, integer, string, boolean, etc
		estos los conocesmos todos pero Golang posee tipos especificos, integers de diferentes bytes
		(8,16,32,64), que se definen de la siguiente manera
	*/
	var x0 int //Depende del sistema operativo, puede ser 32 o 64 bits
	var x1 int8
	var x2 int16
	var x3 int32
	var x4 int64

	/*
		Existen otros tipo de integers los cuales no almacenan signo, estos integer se llaman "uint"
		y con su respectivo valor de byte
	*/
	var x5 uint //Depende del sistema operativo, puede ser 32 o 64 bits
	var x6 uint8
	var x7 uint16
	var x8 uint32
	var x9 uint64
	/*
		Puntos flotantes hacen referencia a valores numericos decimales, estos pueden ser de 32 o 64 bits
		y pueden ser resultantes de divisiones o multiplicaciones de integer y un punto flotante
	*/
	var x10 float32
	var x11 float64
	/*
		Cadenas de texto
	*/
	var x12 string
	/*
		Booleanos
	*/
	var x13 bool
	/*
		Byte, alternativa de uint8, siendo lo mismo, como un alias
	*/
	var x14 byte
	/*
		Rune, alternativa a int32
	*/
	var x15 rune
}
