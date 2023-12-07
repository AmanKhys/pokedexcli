package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	fmt.Println("Hello World!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >>")
		scanner.Scan()
		commandString := scanner.Text()
		fmt.Printf("the command %v is used \n", commandString)
		fmt.Println()
		command := commandProvider()[commandString]
		err := command.callback()

		fmt.Println()
		fmt.Println()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func commandProvider() map[string]cliCommand {
	//command callback functions, will be used as closures in the main func.
	commandHelp := func() error {
		fmt.Println("Welcome to Pokedex!")
		fmt.Println("Usage:")
		commands := commandProvider()
		for _, value := range commands {
			fmt.Printf("%v: %v \n", value.name, value.description)
		}
		return nil

	}

	commandExit := func() error {
		os.Exit(0)
		return nil
	}
	//returning the commands as map
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
