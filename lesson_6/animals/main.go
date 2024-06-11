package main

import (
	"fmt"
	"os"

	"animals/pets"
)

func main() {
	myCat := pets.Cat{
		Animal:   pets.Animal{Name: "mr. buttons"},
		Age:      5,
		IsAsleep: true,
	}
	myDog := pets.Dog{
		Animal: pets.Animal{Name: "spot"},
		Age:    6,
		Weight: 30,
	}

	var feedToCat uint8 = 3
	var feedToDog uint8 = 10

	catFed, err := feed(&myCat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFed)
	}

	fmt.Print("\n\n\t =====\n\n\n")

	dogFed, err := feed(&myDog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: %v\n", err)
	} else {
		fmt.Println("Dog ate:", dogFed)
	}

	fmt.Print("\n\n\t =====\n\n\n")

	displayInfo(myCat)
	displayInfo(myDog)
	displayInfo(42)
}

func feed(animal pets.EaterWalker, amount uint8) (uint8, error) {
	switch v := animal.(type) {
	case *pets.Cat:
		fmt.Println(v.GetName(), "is a cat aged", v.Age)
	case *pets.Dog:
		fmt.Println(v.GetName(), "is a dog aged", v.Age)
	default:
		fmt.Println("Unknown animal type")
	}

	fmt.Println("First, let's walk!")
	fmt.Println(animal.Walk())

	fmt.Println("Now, let's feed our", animal.GetName())

	return animal.Eat(amount)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is an int:", v)
	case pets.Cat:
		fmt.Println("This is a Cat named:", v.Name, "and it is", v.Age, "years old")
	case pets.Dog:
		fmt.Println("This is a Dog named:", v.Name, "and weight:", v.Weight)
	default:
		fmt.Println("Unknown type")
	}
}
