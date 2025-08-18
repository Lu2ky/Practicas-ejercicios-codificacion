package main

import "fmt"

func main(){
	var s string
	var c int
	fmt.Print("enter a string: ")
	fmt.Scan(&s)
	for i:= 0; i < len(s) ; i++ {
		if (s[i] == 'i' || s[i] == 'I') && (c == 0 && i == 0) {
			c++
		}
		if(s[i] == 'a' || s[i] == 'A') && c == 1{
			c++
		}
		if(s[i] == 'n' || s[i] == 'N') && c == 2{
			c++
		}
	}
	if(c == 3){
		fmt.Println("IAN FOUND!")
	} else {
		fmt.Println("Ian no found :c")
	}
}