package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func CheckCharacter(cfgApiClient *models.Config, character string) error {
	characterName := character

	client := characters.ClientWrapper{Client: cfgApiClient.ApiClient}

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

	return nil
}
