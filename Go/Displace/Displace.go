package main

import (
	"fmt"
)

func main() {
	var t, a, vo, so float64
	fmt.Print("Introduce a (acceleration): ")
	fmt.Scan(&a)
	fmt.Print("Introduce vo (Initial velocity): ")
	fmt.Scan(&vo)
	fmt.Print("Introduce so (Initial displace): ")
	fmt.Scan(&so)
	fn := GenDisplaceFn(a, vo, so)
	fmt.Print("Introduce t (time): ")
	fmt.Scan(&t)
	fmt.Print("Displacement: ", fn(t))

}
func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	function := func(t float64) float64 {
		return (0.5*(a)*t*t + vo*t + so)
	}
	return function
}
