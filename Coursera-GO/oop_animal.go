package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {

	return a.food

}
func (a *Animal) Move() string {
	return a.locomotion

}

func (a *Animal) Speak() string {

	return a.noise

}
func main() {

	for {

		var name string
		var action string

		i_cow := Animal{"grass", "walk", "moo"}
		i_bird := Animal{"worms", "fly", "peep"}
		i_snake := Animal{"mice", "slither", "hsss"}

		fmt.Print(">")

		fmt.Scan(&name)
		fmt.Scan(&action)

		var inputname Animal
		switch {
		case name == "cow":
			inputname = i_cow
		case name == "bird":
			inputname = i_bird
		case name == "snake":
			inputname = i_snake
		}

		switch {
		case action == "eat":
			fmt.Println(inputname.Eat())
		case action == "move":
			fmt.Println(inputname.Move())
		case action == "speak":
			fmt.Println(inputname.Speak())

		}

	}

}
