package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandCreateCharacter(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	characterName := args[0]
	characterSkin := args[1]

	client := cfg.ApiClient
	wrapper := characters.ClientWrapper{Client: client}

	request := models.CharacterCreateRequest{
		Name: characterName,
		Skin: characterSkin,
	}

	character, err := wrapper.Create(request)
	if err != nil {
		return fmt.Errorf("failed to create character: %w", err)
	}

	fmt.Printf("created character: %s\n", character.Data.Name)
	return nil
}
