package main

import (
	"fmt"
)

func Bubblesort(arraytosort []int) {

	for i := 0; i < len(arraytosort)-1; i++ {
		for j := 0; j < len(arraytosort)-i-1; j++ {
			pivot1 := arraytosort[j]
			pivot2 := arraytosort[j+1]
			if pivot2 < pivot1 {
				swap(j, arraytosort)
			}
		}
	}
	for _, z := range arraytosort {
		print(z, " ")
	}
}
func swap(i int, array []int) {
	array[i], array[i+1] = array[i+1], array[i]
}

func main() {
	var originalarray [10]int
	for i := 0; i < 10; i++ {
		fmt.Print("ingrese numero(", i+1, "): ")
		fmt.Scan(&originalarray[i])
	}
	for _, z := range originalarray {
		print(z, " ")
	}
	fmt.Println()
	Bubblesort(originalarray[0:10])
}
