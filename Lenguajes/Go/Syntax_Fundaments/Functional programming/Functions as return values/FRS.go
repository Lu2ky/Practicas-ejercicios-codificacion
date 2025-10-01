package main

import (
	"fmt"
	"math"
)

func main() {
	var o_x, o_y float64
	var x, y float64
	fmt.Print("Ingrese el punto x de origen: ")
	fmt.Scan(&o_x)
	fmt.Println()
	fmt.Print("Ingrese el punto y de origen: ")
	fmt.Scan(&o_y)
	fmt.Print("Ingrese el punto x a calcular distancia: ")
	fmt.Scan(&x)
	fmt.Println()
	fmt.Print("Ingrese el punto y a calcular distancia: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Println()
	fmt.Print(MakeDistOrigin(o_x, o_y)(x, y))
}
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {

	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}
