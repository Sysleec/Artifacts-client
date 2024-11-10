package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
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
		go func() {
			client := artsapi.NewClient(5*time.Second, cfg.ApiClient.Token)
			newCfg := models.Config{
				ApiClient: &client,
				DB:        cfg.DB,
			}

			err := commandBotStart(&newCfg, character.Name, "fighting", "chicken")
			if err != nil {
				fmt.Printf("failed to start bot for character %s: %v\n", character.Name, err)
			}
		}()

	}

	fmt.Println("Auto bot for all characters started")

	return nil
}
