package data

import "math"

type Point struct {
	X float64
	Y float64
}

func (p Point) DistToOrig() float64 {
	t := math.Pow(p.X, 2) + math.Pow(p.Y, 2)
	return math.Sqrt(t)
}
