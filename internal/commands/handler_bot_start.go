package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/bot"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func commandBotStart(cfg *models.Config, args ...string) error {
	if len(args) != 3 {
		return fmt.Errorf("expected exactly 3 arguments, got %d", len(args))
	}

	character := args[0]

	if cfg.ApiClient.BotRunning[character] == true {
		return fmt.Errorf("bot for character %s is already running", character)
	}

	client, err := bot.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	char, err := utils.CheckCharacter(cfg, character)
	if err != nil {
		return fmt.Errorf("failed to check character: %w", err)
	}

	isMaxItems := utils.CheckMaxItems(char)
	if isMaxItems {
		return fmt.Errorf("character %s has max items", character)
	}

	switch args[1] {
	case "gather":
		switch args[2] {
		case "copper":
			err := client.Gather("copper", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "gudgeon":
			err := client.Gather("gudgeon", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "iron":
			err := client.Gather("iron", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		default:
			return fmt.Errorf("unknown target: %s", args[2])
		}
	case "fighting":
		switch args[2] {
		case "chicken":
			err := client.Fighting("chicken", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "cow":
			err := client.Fighting("cow", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "green_slime":
			err := client.Fighting("green_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "blue_slime":
			err := client.Fighting("blue_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "yellow_slime":
			err := client.Fighting("yellow_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "red_slime":
			err := client.Fighting("red_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		}
	default:
		return fmt.Errorf("unknown action: %s", args[1])
	}

	fmt.Printf("Bot for character %s started\n", character)

	return nil
}
