package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/bot"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandBotStop(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	character := args[0]

	if cfg.ApiClient.BotRunning[character] == false {
		return fmt.Errorf("bot for character %s is not running", character)
	}

	client, err := bot.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	client.Client.BotRunning[character] = false

	fmt.Printf("Bot for character %s stopping...\n", character)

	return nil
}
