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

	err = utils.CheckMaxItems(char)
	if err != nil {
		return err
	}

	switch args[1] {
	case "mining":
		switch args[2] {
		case "copper":
			err := client.MiningCopper(character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		default:
			return fmt.Errorf("unknown target: %s", args[2])
		}
	default:
		return fmt.Errorf("unknown action: %s", args[1])
	}

	fmt.Printf("Bot for character %s started\n", character)

	return nil
}
