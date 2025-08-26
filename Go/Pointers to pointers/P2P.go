package main

import "fmt"

func main() {
	var pminus1 int
	var p0 ****int
	var p1 ***int
	var p2 **int
	var p3 *int
	var p4 int
	pminus1 = 1
	p3 = &pminus1
	p2 = &p3
	p1 = &p2
	p0 = &p1
	p4 = ****p0
	fmt.Println(p4)
}
