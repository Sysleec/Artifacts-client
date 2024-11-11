package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/characters"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func commandBotAutoStart(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

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

	action, resourse := args[0], args[1]

	for _, character := range myCharacters.Data {
		go func() {
			client := artsapi.NewClient(5*time.Second, cfg.ApiClient.Token)
			newCfg := models.Config{
				ApiClient: &client,
				DB:        cfg.DB,
			}

			err := commandBotStart(&newCfg, character.Name, action, resourse)
			if err != nil {
				fmt.Printf("failed to start bot for character %s: %v\n", character.Name, err)
			}
		}()

	}

	fmt.Println("Auto bot for all characters started")

	return nil
}
