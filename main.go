package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {	
	return map[string]cliCommand {
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		},
		"help": {
			name:		 "help",
			description: "List of commands for Pokedex",
			callback:	 commandHelp,
		},
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

func main() {
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
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Printf("Unknown command: %s. Type 'help' to see available commands.\n", command)
				continue
			}
		} else {
			fmt.Println("Please enter a command. You can type 'help' to find out how to use Pokedex.")
		}
	}
}
