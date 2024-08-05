package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandGetCharacter(cfg *models.Config, args ...string) error {
	characterName := args[0]

	client := cfg.ApiClient
	wrapper := characters.ClientWrapper{Client: client}

	character, err := wrapper.Get(characterName)
	if err != nil {
		return fmt.Errorf("failed to create character: %w", err)
	}

	fmt.Printf("get character: %s\n", character.Data.Name)
	return nil
}
