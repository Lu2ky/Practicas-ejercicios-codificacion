package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre string
	Edad   int
	Id     int
}

func main() {
	/*
		Los archivos json o Java Script Object Notation,
		nos sirven para comunicar sistemas entre si, usado en mayor parte
		para la devolucion de peticios de apis

		json.Marchal(T) --> Convierte el dato en formato JSON []Byte, tiene 2 salidas, el JSON y el error, en
							caso de que no se pueda convertir a JSON,
							JSON, ERR := json.Marshal(T)
		json.Unmarshal(JSON, T1) --> Devuelve el objeto si cuadra con el molde de GoLang
									 los argumentos son, JSON a devolver a objeto de GoLang
									 T1 es el puntero de la variable donde se va a almacenar
									 el resultado, el metodo solo tiene una salida, la de error

		El "JSON" del comando "json.Unmarshal" es el Jso

	*/
	p1 := persona{Nombre: "Juan", Edad: 18, Id: 9123}

	bar, _ := json.Marshal(p1)
	fmt.Println(bar)

	var p2 persona

	json.Unmarshal(bar, &p2)

	fmt.Println(p2)
}
