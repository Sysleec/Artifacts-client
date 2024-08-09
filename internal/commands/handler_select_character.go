package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandSelectCharacter(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	characterName := args[0]

	client := characters.ClientWrapper{Client: cfg.ApiClient}

	myCharacters, err := client.GetMyCharacters()
	if err != nil {
		return fmt.Errorf("failed to get my myCharacters: %w", err)
	}

	found := false

	for _, character := range myCharacters.Data {
		if character.Name == characterName {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("your character not found: %s", characterName)
	}

	client.Client.Character = characterName

	fmt.Printf("selected character: %s\n", client.Client.Character)
	return nil
}
