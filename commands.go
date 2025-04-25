package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func GetCommands() map[string]cliCommand {
	commands := make(map[string]cliCommand, 2)
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a message",
		callback:    CliHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the pokedex",
		callback:    CliExit,
	}
	return commands
}

func CliHelp() error {
	message := "Welcome to the pokedex!\nUsage:\n\n"
	for k, v := range GetCommands() {
		message += k + ": " + v.description + "\n"
	}
	_, err := fmt.Println(message)
	if err != nil {
		return fmt.Errorf("error printing to the console: %w", err)
	}
	return nil
}

func CliExit() error {
	os.Exit(0)
	return nil
}
