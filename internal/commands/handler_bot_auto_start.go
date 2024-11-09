package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandBotAutoStart(cfg *models.Config, _ ...string) error {
	clientWrapper, err := characters.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to create client wrapper: %w", err)
	}

	myCharacters, err := clientWrapper.GetMyCharacters()
	if err != nil {
		return fmt.Errorf("failed to get my characters: %w", err)
	}

	if len(myCharacters.Data) == 0 {
		return fmt.Errorf("no characters found")
	}

	for _, character := range myCharacters.Data {
		err := commandBotStart(cfg, character.Name, "gather", "copper")
		if err != nil {
			return fmt.Errorf("failed to start bot for character %s: %w", character.Name, err)
		}
	}

	fmt.Println("Auto bot for all characters started")

	return nil
}
