package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"strconv"
	"strings"
	"time"
)

func commandCraft(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	target := strings.ToLower(args[0])

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	qty, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("qty must be an integer: %w", err)
	}

	request := models.CraftReq{
		Code:     target,
		Quantity: qty,
	}

	character, err := client.Craft(request)
	if err != nil {
		return fmt.Errorf("failed to craft: %w", err)
	}

	fmt.Printf("Crafted %d %s\n", qty, target)

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