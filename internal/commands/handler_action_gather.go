package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func commandGather(cfg *models.Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected exactly 0 arguments, got %d", len(args))
	}

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	character, err := client.Gathering()
	if err != nil {
		return fmt.Errorf("failed to gather resources: %w", err)
	}

	fmt.Printf("Gathered resources at x = %d, y = %d\n", character.Data.Destination.X, character.Data.Destination.Y)

	secondsRemaining := character.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	isMaxItems := utils.CheckMaxItems(models.ConvertToModelCharacter(character))
	if isMaxItems {
		return fmt.Errorf("character %s has max items", character)
	}

	return nil
}
