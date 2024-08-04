package commands

import "github.com/Sysleec/Artifacts-client/internal/models"

type cliCommand struct {
	name        string
	description string
	Callback    func(*models.Config, ...string) error
}

func List() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Artifacts client",
			Callback:    commandExit,
		},
	}
}