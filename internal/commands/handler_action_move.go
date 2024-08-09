package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"strconv"
)

func commandMove(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	moveX := args[0]
	moveY := args[1]

	client := action.ClientWrapper{Client: cfg.ApiClient}

	moveXInt, err := strconv.Atoi(moveX)
	if err != nil {
		return fmt.Errorf("x must be an integer: %w", err)
	}

	moveYInt, err := strconv.Atoi(moveY)
	if err != nil {
		return fmt.Errorf("y must be an integer: %w", err)
	}

	request := models.MoveReq{
		X: moveXInt,
		Y: moveYInt,
	}

	character, err := client.Move(request)
	if err != nil {
		return fmt.Errorf("failed to move character: %w", err)
	}

	fmt.Printf("Moved character to x = %d, y = %d\n", character.Data.Destination.X, character.Data.Destination.Y)

	return nil
}
