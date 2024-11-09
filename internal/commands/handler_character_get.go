package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
)

func commandGetCharacter(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly  argument, got %d", len(args))
	}

	characterName := args[0]

	client := characters.ClientWrapper{Client: cfg.ApiClient}

	character, err := client.Get(characterName)
	if err != nil {
		return fmt.Errorf("failed to create character: %w", err)
	}

	utils.CharacterPrettyPrinter(character.Data)

	return nil
}
