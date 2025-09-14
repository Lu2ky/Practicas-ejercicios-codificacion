package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Nombre struct {
	Fname string
	Lname string
}

func main() {
	var filename string
	fmt.Print("Ingrese el nombre del archivo: ")
	fmt.Scanln(&filename)
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error")
		return
	}
	lineas := strings.Split(string(dat), "\n")
	var save []Nombre
	for _, dato := range lineas {
		datolinea := strings.SplitN(dato, " ", 2)
		if dato == "" {
			continue
		}
		if len(datolinea[0]) < 20 && len(datolinea[1]) < 20 {
			var temp Nombre
			temp.Fname = datolinea[0]
			temp.Lname = datolinea[1]
			save = append(save, temp)
		} else {
			continue
		}
	}
	for _, structdata := range save {
		fmt.Println("FName: " + structdata.Fname + " LName: " + structdata.Lname)
	}
}
