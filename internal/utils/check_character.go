package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func CheckCharacter(cfgApiClient *models.Config, character string) (models.Character, error) {
	characterName := character

	client := characters.ClientWrapper{Client: cfgApiClient.ApiClient}

	myCharacters, err := client.GetMyCharacters()
	if err != nil {
		return models.Character{}, fmt.Errorf("failed to get my myCharacters: %w", err)
	}

	found := false

	var char models.Character

	for _, character := range myCharacters.Data {
		if character.Name == characterName {
			found = true
			char = character
			break
		}
	}

	if !found {
		return models.Character{}, fmt.Errorf("your character not found: %s", characterName)
	}

	return char, nil
}
