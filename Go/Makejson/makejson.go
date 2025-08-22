package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Print("Enter a name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter a address: ")
	fmt.Scanln(&address)
	Map := map[string]string{
		"name":    name,
		"address": address,
	}
	Json, _ := json.Marshal(Map)
	fmt.Println(string(Json))

}
