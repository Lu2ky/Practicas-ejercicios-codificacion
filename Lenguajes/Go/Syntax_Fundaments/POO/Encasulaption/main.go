package main

import (
	"Encasulaption/data"
	"fmt"
)

func main() {
	var p data.Point
	p.X = 3
	p.Y = 4
	fmt.Print(p.DistToOrig())
}
