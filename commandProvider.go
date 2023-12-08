package main

func commandProvider() map[string]cliCommand {
	//command callback functions, will be used as closures in the main func.

	urlString = "https://pokeapi.co/api/v2/location-area/"

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
		"map": {
			name:        "map",
			description: "prints next twenty locations",
			callback:    commandMap,
		},
		"bmap": {
			name:        "bmap",
			description: "prints previous twenty locations",
			//			callback:    commandBmap,
		},
	}
}
