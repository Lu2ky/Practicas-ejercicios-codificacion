package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		La libreria ioutil sirve para interactuar con
		archivos, no solo de texto, aqui solo abarcaremos estos.

		El primer metodo es .ReadFile("filename")
		el cual nos ayuda a leer un archivo, este nos devuelve
		dos cosas datos en []byte y un erro
	*/
	dat, _ := ioutil.ReadFile("test.txt")
	datS := string(dat)
	/*
		Con el metodo string(var) pasamos de []byte a string
		ya teniendo los datos a un formato legible para el ser humano
		y los imprimimos
	*/
	fmt.Println(datS)

	dat1 := "Hello, world!"
	dat1b := []byte(dat1)
	/*
		Aqui hacemos un proceso inverso, convertimos un char,
		luego mandamos este dat1b a un nuevo metodo llamado
		.WriteFile el cual recibe 3 argumentos, nombre del archivo,
		texto en []byte, y los permisos asociados, estos permisos deben
		tener un prefijo con "0" debido a que lo estamos representando en
		base 10, el sistema de permisos de los archivos esta en base 8

		TABLA DE PERMISOS

		Número | Binario | Lectura (r) | Escritura (w) | Ejecución (x)
		---------------------------------------------------------------
		0      | 000     |     -       |      -        |      -
		1      | 001     |     -       |      -        |      x
		2      | 010     |     -       |      w        |      -
		3      | 011     |     -       |      w        |      x
		4      | 100     |     r       |      -        |      -
		5      | 101     |     r       |      -        |      x
		6      | 110     |     r       |      w        |      -
		7      | 111     |     r       |      w        |      x

		O ver el siguiente canva, sección de notacion numerica:
		https://www.canva.com/design/DAGwpR_UrC8/HpK21PZAv01DPv6sSb-dxg/edit?utm_content=DAGwpR_UrC8&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton

	*/
	ioutil.WriteFile("newtest.txt", dat1b, 0777)
}
