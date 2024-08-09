package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandDeleteCharacter(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	characterName := args[0]

	client := characters.ClientWrapper{Client: cfg.ApiClient}

	request := models.CharacterDeleteRequest{
		Name: characterName,
	}

	character, err := client.Delete(request)
	if err != nil {
		return fmt.Errorf("failed to delete character: %w", err)
	}

	fmt.Printf("deleted character: %s\n", character.Data.Name)
	return nil
}
