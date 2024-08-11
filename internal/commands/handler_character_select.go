package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func commandSelectCharacter(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	character := args[0]

	_, err := utils.CheckCharacter(cfg, character)
	if err != nil {
		return fmt.Errorf("failed to check character: %w", err)
	}

	client := characters.ClientWrapper{Client: cfg.ApiClient}

	client.Client.Character = character

	fmt.Printf("selected character: %s\n", client.Client.Character)
	return nil
}
