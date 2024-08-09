package commands

import "github.com/Sysleec/Artifacts-client/internal/models"

type cliCommand struct {
	name        string
	description string
	Callback    func(*models.Config, ...string) error
}

func List() map[string]cliCommand {
	return map[string]cliCommand{
		"move": {
			name:        "move <x> <y>",
			description: "Move the selected character to the specified coordinates",
			Callback:    commandMove,
		},
		"character create": {
			name:        "character create <name> <skin>",
			description: "Create a new character",
			Callback:    commandCreateCharacter,
		},
		"character delete": {
			name:        "character delete <name>",
			description: "Delete a character by name",
			Callback:    commandDeleteCharacter,
		},
		"character get": {
			name:        "character get <name>",
			description: "Get a character by name",
			Callback:    commandGetCharacter,
		},
		"character select": {
			name:        "character select <name>",
			description: "Select a character by name for actions",
			Callback:    commandSelectCharacter,
		},
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
