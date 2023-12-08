package main

import (
	"bufio"
	//	"encoding/json"
	"fmt"
	//	"io"
	//	"log"
	//	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
		command, ok := commandProvider()[commandString]
		if ok == false {
			fmt.Printf("%v is not valid, enter a valid command \n", commandString)
			continue
		}
		err := command.callback()

		fmt.Println()
		fmt.Println()
		if err != nil {
			fmt.Println(err)
		}
	}

}
