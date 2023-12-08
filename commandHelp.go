package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	commands := commandProvider()
	for _, value := range commands {
		fmt.Printf("%v: %v \n", value.name, value.description)
	}
	return nil

}
