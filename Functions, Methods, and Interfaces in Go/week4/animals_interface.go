/* Write a program which allows the user to create a set of animals and to get information about those animals.
 * Each animal has a name and can be either a cow, bird, or snake. 
 * With each command, the user can either create a new animal of one of the three types, 
 * or the user can request information about an animal that he/she has already created. 
 * Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, 
 * but the types of animals are restricted to either cow, bird, or snake. 
 * The following table contains the three types of animals and their associated data.

 * Animal	Food eaten	Locomotion method	Spoken sound
 * cow	        grass	        walk	           	moo
 * bird	        worms	        fly	                peep
 * snake	    mice	        slither	            hsss
 
 * Your program should present the user with a prompt, “>”, to indicate that the user can type a request. 
 * Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. 
 * Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.
 
 * Each “newanimal” command must be a single line containing three strings. 
 * The first string is “newanimal”. 
 * The second string is an arbitrary string which will be the name of the new animal. 
 * The third string is the type of the new animal, either “cow”, “bird”, or “snake”. 
 * Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

 * Each “query” command must be a single line containing 3 strings. 
 * The first string is “query”. 
 * The second string is the name of the animal. 
 * The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
 * Your program should process each query command by printing out the requested data.
 
 * Define an interface type called Animal which describes the methods of an animal. 
 * Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. 
 * The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, 
 * and the Speak() method should print the animal’s spoken sound. 
 * Define three types Cow, Bird, and Snake. 
 * For each of these three types, define methods Eat(), Move(), and Speak() 
 * so that the types Cow, Bird, and Snake all satisfy the Animal interface. 
 * When the user creates an animal, create an object of the appropriate type. 
 * Your program should call the appropriate method when the user issues a query command.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// A map that key is animal name and value is animal type
	my_map := new(AllMap)
	my_map.cow = make(map[string]Cow)
	my_map.bird = make(map[string]Bird)
	my_map.snake = make(map[string]Snake)
	my_map.animal = make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(">")
		ipt_raw, _ := reader.ReadString('\n')
		ipt := strings.TrimSuffix(ipt_raw, "\n")
		sli_str := strings.Split(ipt, " ")
		if len(sli_str) != 3 {
			fmt.Println("Illegal input, please try again.")
			continue
		}
		command := sli_str[0]
		arg_1 := sli_str[1]
		arg_2 := sli_str[2]

		switch command {
		case "newanimal":
			NewAnimal(arg_1, arg_2, my_map)
		case "query":
			Query(arg_1, arg_2, my_map)
		default:
			fmt.Printf("Command \"%v\" not found. Please try again\n", command)
		}
	}

}

func NewAnimal(a_name string, a_type string, my_map *AllMap) {
	//new_animal := new(Animal)
	switch a_type {
	case "cow":
		my_map.animal[a_name] = "cow"
		my_map.cow[a_name] = *new(Cow)
	case "bird":
		my_map.animal[a_name] = "bird"
		my_map.bird[a_name] = *new(Bird)
	case "snake":
		my_map.animal[a_name] = "snake"
		my_map.snake[a_name] = *new(Snake)
	default:
		fmt.Printf("Animal type \"%v\" not found. Please try again\n", a_type)
	}

}

/* func CreateNewAnimal(a Animal, a_name string, animal_map map[string]string){
	animal_map[a_name] = a
	fmt.Println("Created it!")
} */

func Query(a_name string, info string, my_map *AllMap) {
	if a_type, ok := my_map.animal[a_name]; ok {
		switch a_type {
		case "cow":
			GetInfo(my_map.cow[a_name], info)	
		case "bird":
			GetInfo(my_map.bird[a_name], info)	
		case "snake":
			GetInfo(my_map.snake[a_name], info)	
		}
	} else {
		fmt.Printf("Animal name \"%v\" not found. Please try again\n", a_name)
	}
}


func GetInfo(a Animal, info string) {
    switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Printf("Information \"%v\" not found. Please try again\n", info)
	}
}

type AllMap struct {
	cow map[string]Cow
	bird map[string]Bird
	snake map[string]Snake
	animal map[string]string
}


type Animal interface {
	Eat() 
	Move()
	Speak()
}

type Cow struct {

}

type Bird struct {
	
}

type Snake struct {

}

func (c Cow) Eat() {
	fmt.Println("grass")
}


func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}
