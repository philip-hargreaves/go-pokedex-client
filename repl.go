package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Command not found")
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	formattedText := strings.TrimSpace(strings.ToLower(text))
	return strings.Fields(formattedText)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
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
			description: "Lists next batch of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous batch of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "View details about a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
