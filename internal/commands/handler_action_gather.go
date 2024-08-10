package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
)

func commandGather(cfg *models.Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected exactly 0 arguments, got %d", len(args))
	}

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	character, err := client.Gathering()
	if err != nil {
		return fmt.Errorf("failed to move character: %w", err)
	}

	fmt.Printf("Gathered resources at x = %d, y = %d\n", character.Data.Destination.X, character.Data.Destination.Y)

	return nil
}
