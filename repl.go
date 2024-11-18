package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	callbackHelp(cfg)

	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		commandName, ok := availableCommands[command]
		if !ok {
			fmt.Println("Invalid Command!")
			continue
		}
		err := commandName.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}

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
			description: "prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List some location ares",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List some location ares (map backwards)",
			callback:    callbackMapb,
		},
		// "explore": {
		// 	name:        "explore",
		// 	description: "List all pokemon in a area",
		// 	callback:    callbackExpore,
		// },
		// "catch": {
		// 	name:        "catch {pokemon name}",
		// 	description: "Attempt to catch pokemon and add it to our pokedex ",
		// 	callback:    callbackCatch,
		// },
		"exit": {
			name:        "exit",
			description: "exists from the cli",
			callback:    callbackExit,
		},
		// "pokedex": {
		// 	name:        "pokedex",
		// 	description: "View all pokemon in pokedex",
		// 	callback:    callbackPokedex,
		// },
		// "inspect": {
		// 	name:        "inspect {pokemon-name}",
		// 	description: "Inspects the caught pokemon stats",
		// 	callback:    callbackInspect,
		// },
	}
}

func cleanInput(str string) []string {
	lowStr := strings.ToLower(str)
	words := strings.Fields(lowStr)
	return words
}
