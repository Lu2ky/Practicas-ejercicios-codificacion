package main

import (
	"fmt"
	"math"
)

type point struct {
	X float64
	Y float64
}

func (p point) DistToOrig() float64 {
	t := math.Pow(p.X, 2) + math.Pow(p.Y, 2)
	return math.Sqrt(t)
}
func main() {
	p1 := point{X: 3, Y: 2}
	fmt.Print(p1.DistToOrig())
}
