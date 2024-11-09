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
		"sell all": {
			name:        "sell all",
			description: "Sell all items",
			Callback:    commandSellAll,
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
		"account set": {
			name:        "account set {name}",
			description: "Set the active account",
			Callback:    commandAccountSet,
		},
		"account list": {
			name:        "account list",
			description: "List all accounts",
			Callback:    commandAccountList,
		},
		"account detail": {
			name:        "account detail",
			description: "Show account details",
			Callback:    commandAccountDetail,
		},
		"resources upload": {
			name:        "resources upload",
			description: "Upload all resources from API artifactsmmo",
			Callback:    commandResourcesUpload,
		},
		"maps upload": {
			name:        "maps upload",
			description: "Upload all maps from API artifactsmmo",
			Callback:    commandMapUpload,
		},
		"maps get": {
			name:        "maps get",
			description: "Get all maps from the database",
			Callback:    commandMapGet,
		},
		"bot auto start": {
			name:        "bot auto start",
			description: "Start automatic actions for all characters",
			Callback:    commandBotAutoStart,
		},
	}
}
