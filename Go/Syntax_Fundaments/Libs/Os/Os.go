package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		La librería os nos permite trabajar con archivos de manera más directa.

		El primer método que usaremos es .Open("filename"),
		el cual recibe como argumento el nombre del archivo (string) que queremos abrir.
		Este método devuelve dos valores:
			- un puntero al archivo abierto
			- un error en caso de que no se haya podido abrir
	*/
	f, _ := os.Open("example.txt")

	/*
		Para leer el contenido del archivo usamos un slice de tipo []byte,
		que nos servirá como "buffer" para almacenar lo leído.

		En este caso usamos make([]byte, 10) para crear un slice de 10 bytes,
		esto quiere decir que leeremos los primeros 10 caracteres del archivo.
	*/
	CharByte := make([]byte, 10)
	f.Read(CharByte)

	/*
		El método .Read() guarda los datos en el slice []byte.
		Para hacerlo legible, convertimos de []byte a string usando string().
	*/
	CharByteToString := string(CharByte)
	fmt.Println(CharByteToString)

	/*
		Siempre que abrimos un archivo con .Open(), debemos cerrarlo al final
		con .Close() para liberar recursos.
	*/
	f.Close()

	/*
		Ahora usamos .Create("filename"), que crea un nuevo archivo.
		Si el archivo ya existe, lo sobreescribe.
		Este método devuelve un puntero al archivo creado y un error si falla.
	*/
	f1, _ := os.Create("OutputExample.txt")

	/*
		El método .Write([]byte) nos permite escribir directamente bytes en el archivo.
		Aquí escribimos un arreglo []byte con tres valores (1,2,3).
		⚠️ Estos no son caracteres visibles, son valores binarios.
	*/
	CharByte2 := []byte{1, 2, 3}
	f1.Write(CharByte2)

	/*
		También podemos escribir directamente texto con .WriteString("texto"),
		lo cual es más cómodo que armar []byte si lo que queremos es texto plano.
	*/
	f1.WriteString("HelloWorld")
}
