package main

import "fmt"

var anm Animal

type Animal struct {
	food       string
	noise      string
	locomotion string
}

func (anm Animal) Eat() {
	fmt.Println(anm.food)
}

func (anm Animal) Move() {
	fmt.Println(anm.locomotion)
}

func (anm Animal) Speak() {
	fmt.Println(anm.noise)
}

func main() {

	var animalName, do string

	for true {
		animalName, do = "", ""
		// Inputs
		fmt.Print("> ")
		_, err := fmt.Scanf("%s %s", &animalName, &do)

		if err != nil {
			fmt.Println(" ")
		}

		// Initialize

		switch animalName {
		case "cow":
			{
				anm = Animal{
					food:       "grass",
					noise:      "moo",
					locomotion: "walk"}
			}
		case "bird":
			{
				anm = Animal{
					food:       "worms",
					noise:      "peep",
					locomotion: "fly"}
			}
		case "snake":
			{
				anm = Animal{
					food:       "mice",
					noise:      "hsss",
					locomotion: "slither"}
			}
		default:
			{
				anm = Animal{
					food:       "mice",
					noise:      "hsss",
					locomotion: "slither"}
			}

		}

		// Do smth
		switch do {
		case "eat":
			anm.Eat()
		case "move":
			anm.Move()
		case "speak":
			anm.Speak()

		}
	}
}
