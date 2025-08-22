package main

import "fmt"
/*
Escriba un programa que pida al usuario que introduzca una cadena. 
El programa busca en la cadena introducida los caracteres 'i', 'a' y 'n'. 
El programa debe imprimir "¡Encontrado!" si la cadena introducida comienza con el carácter 'i', termina con el carácter 'n' y contiene el carácter 'a'. 
El programa debe imprimir "¡No encontrado!" en caso contrario. El programa no debe distinguir entre mayúsculas y minúsculas, por lo que no importa si los caracteres son mayúsculas o minúsculas.
Ejemplos: El programa debe imprimir "¡Encontrado!" para las siguientes cadenas introducidas de ejemplo, "ian", "Ian", "iuiygaygn", "I d skd a efju N". 
El programa debe imprimir "¡No encontrado!" para las siguientes cadenas, "ihhhhhn", "ina", "xian". 
*/
func main(){
	var s string
	var c int
	fmt.Print("enter a string: ")
	fmt.Scan(&s)
	for i:= 0; i < len(s) ; i++ {
		if (s[i] == 'i' || s[i] == 'I') && (c == 0 && i == 0) {
			c++
		}
		if(s[i] == 'a' || s[i] == 'A') && c == 1{
			c++
		}
		if(s[i] == 'n' || s[i] == 'N') && c == 2{
			c++
		}
	}
	if(c == 3){
		fmt.Println("IAN FOUND!")
	} else {
		fmt.Println("Ian no found :c")
	}
}
