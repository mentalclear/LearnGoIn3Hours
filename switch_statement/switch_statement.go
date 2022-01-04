package main

import (
	"fmt"
	"os"
)

func lotsOfIfs() {
	word := os.Args[1]
	if word == "hello" {
		fmt.Println("Hi yourself")
	} else if word == "goodbye" {
		fmt.Println("So long!")
	} else if word == "greetings" {
		fmt.Println("Salutations")
	} else {
		fmt.Println("I don't know what you said")
	}
}

func with_switch() {
	word := os.Args[1]
	greet := "greetings"
	switch l := len(word); word {
	case "hi":
		fmt.Println("Very informal")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself")
	case "goodbye", "bye":
		fmt.Println("So long!")
	case greet:
		fmt.Println("Salutations")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long")
	}
}

func fancy_switch() {
	word := os.Args[1]
	c := "crackerjack"
	switch l := len(word); {
	case word == "hi":
		fmt.Println("Very informal!")
		fallthrough
	case word == "hello":
		fmt.Println("Hi yourself")
	case l == 1:
		fmt.Println("I don't know any one letter words")
	case 1 < l && l < 10, word == c:
		fmt.Println("This word is either", c, "or 2-9 characters long")
	default:
		fmt.Println("I don't know what you said but it was", l, "characters long.")
	}
}

func main() {
	//lotsOfIfs()
	//with_switch()
	fancy_switch()
}
