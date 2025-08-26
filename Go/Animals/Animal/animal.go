package Animal

import "fmt"

type Animal struct {
	Food       string
	Locomotion string
	Spoken     string
}

func (ts Animal) Eat() {
	fmt.Println(ts.Food)
}
func (ts Animal) Move() {
	fmt.Println(ts.Locomotion)
}
func (ts Animal) Speak() {
	fmt.Println(ts.Spoken)
}
