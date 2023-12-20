package main

import "fmt"

/*

=> Coursera.Org <=

Course:
	Functions, Methods, and Interfaces in Go

* Assessments of WEEK 3

##################################
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
	1) the food that it eats,
	2) its method of locomotion,
	3) the sound it makes when it speaks.

The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal  ==   Food eaten  == Locomotion method == Spoken sound

cow     ==     grass     ==        walk       ==     moo

bird    ==     worms     ==        fly        ==     peep

snake   ==     mice      ==      slither      ==     hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal.
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak().
The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method
should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.

****************************************

This assignment is worth a total of 10 points.

Test the program by running it and testing it by issuing three requests.
Each request should involve a different animal (cow, bird, snake) and a different property of the animal (eat, move, speak).
Give 2 points for each request for which the program provides the correct response.

Examine the code. If the code contains a type called Animal,
which is a struct containing three fields, all of which are strings,
then give another 2 points. If the program contains three methods called Eat(), Move(), and Speak(),
and all of their receiver types are Animal, give another 2 points.

##################################
##################################
##################################
##################################

/*

> animalName  eat/move/speak
[print answer]
*/

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

		// Inputs
		fmt.Print("> ")
		_, err := fmt.Scanf("%s %s", &animalName, &do)

		if err != nil {
			fmt.Println("Error !")
			break
		}

		// Initialize

		switch animalName {
		case "cow":
			anm = Animal{
				food:       "grass",
				noise:      "moo",
				locomotion: "walk"}

		case "bird":
			anm = Animal{
				food:       "worms",
				noise:      "peep",
				locomotion: "fly"}

		case "snake":
			anm = Animal{
				food:       "mice",
				noise:      "hsss",
				locomotion: "slither"}
		}

		switch do {
		case "eat":
			//do somth
		case "move":
			//do somth
		case "speak":
		//do somth

	}
}
