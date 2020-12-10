package main

import (
	"fmt"
)

type animalStruct struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (as animalStruct) Eat() {
	fmt.Println(as.food)
	return
}

func (as animalStruct) Move() {
	fmt.Println(as.locomotion)
	return
}

func (as animalStruct) Speak() {
	fmt.Println(as.noise)
	return
}

func main() {
	//ai:  animal interface
	//am:    animal map
	//as:  animal struct
	am := make(map[string]animalStruct)
	am["cow"] = animalStruct{"grass", "walk", "moo"}
	am["bird"] = animalStruct{"worms", "fly", "peep"}
	am["snake"] = animalStruct{"mice", "slither", "hsss"}
	var ai animalInterface
	for {
		var command, requestAni, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			ai = am[requestAni]
			switch requestType {
			case "eat":
				ai.Eat()
			case "move":
				ai.Move()
			case "speak":
				ai.Speak()
			}
		}
		if command == "newanimal" {
			am[requestAni] = am[requestType]
			fmt.Println("Created it!")
		}
	}
}
