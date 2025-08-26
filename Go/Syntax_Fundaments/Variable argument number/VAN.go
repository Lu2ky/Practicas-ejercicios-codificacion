package main

import "fmt"

func main() {
	fmt.Println(Max(2, 3, 5, 1, 2, 4, 12, 5, 76, 12, 768, 12, 43, 12, 21))
}
func Max(values ...int) int {
	max := -1
	for _, z := range values {
		if z > max {
			max = z
		}
	}
	return max
}
