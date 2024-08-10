package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"strconv"
	"time"
)

func commandMove(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	moveX := args[0]
	moveY := args[1]

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

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

	secondsRemaining := character.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	return nil
}
