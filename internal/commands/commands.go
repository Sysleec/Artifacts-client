package commands

import "github.com/Sysleec/Artifacts-client/internal/models"

type cliCommand struct {
	name        string
	description string
	Callback    func(*models.Config, ...string) error
}

func List() map[string]cliCommand {
	return map[string]cliCommand{
		"buy": {
			name:        "buy <item_code> <amount>",
			description: "Buy an item",
			Callback:    commandBuy,
		},
		"fight": {
			name:        "fight",
			description: "Fight a monster",
			Callback:    commandFight,
		},
		"sell": {
			name:        "sell <item_code> <amount>",
			description: "Sell an item",
			Callback:    commandSell,
		},
		"unequip": {
			name:        "unequip <slot>",
			description: "Unequip an item",
			Callback:    commandUnequip,
		},
		"equip": {
			name:        "equip <item_code> <slot>",
			description: "Equip an item",
			Callback:    commandEquip,
		},
		"craft": {
			name:        "craft <item_code> <amount>",
			description: "Craft an item",
			Callback:    commandCraft,
		},
		"bot start": {
			name:        "bot start <character> <action> <target>",
			description: "Bot for automatic actions",
			Callback:    commandBotStart,
		},
		"bot stop": {
			name:        "bot stop <character>",
			description: "Stop the bot",
			Callback:    commandBotStop,
		},
		"gather": {
			name:        "gather",
			description: "Gather resources",
			Callback:    commandGather,
		},
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
