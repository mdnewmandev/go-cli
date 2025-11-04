package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdnewmandev/go-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient 	 pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) > 0 {
			commandName := cleaned[0]
			command, exists := getCommands()[commandName]

			if exists {
				params := cleaned[1:]
				err := command.callback(cfg, params)
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Printf("Unknown command: %s. Type 'help' to see available commands.\n", commandName)
				continue
			}
		} else {
			fmt.Println("Please enter a command. You can type 'help' to find out how to use Pokedex.")
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)
	result := make([]string, 0, len(words))
	for _, word := range words {
		if word != "" {
			result = append(result, word)
		}
	}

	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
	params      []string  // Optional: describes expected parameters
}

func getCommands() map[string]cliCommand {	
	return map[string]cliCommand{
		"help": {
			name:		 "help",
			description: "List of commands for Pokedex",
			callback:	 commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:	 commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:	 commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location by name",
			callback:	 commandExplore,
			params:		 []string{"location_name"},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		},
	}
}