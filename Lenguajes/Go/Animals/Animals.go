package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var animalcow, animalbird, animalsnake Animal
	scanner := bufio.NewScanner(os.Stdin)
	animalcow.Food, animalcow.Locomotion, animalcow.Spoken = "grass", "walk", "moo"
	animalbird.Food, animalbird.Locomotion, animalbird.Spoken = "worms", "fly", "peep"
	animalsnake.Food, animalsnake.Locomotion, animalsnake.Spoken = "mice", "slither", "hsss"
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		desc := scanner.Text()
		desc1 := strings.Split(desc, " ")
		if desc1[0] == "cow" {
			action(animalcow, desc1[1])
		}
		if desc1[0] == "bird" {
			action(animalbird, desc1[1])
		}
		if desc1[0] == "snake" {
			action(animalsnake, desc1[1])
		}
	}
}
func action(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

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
