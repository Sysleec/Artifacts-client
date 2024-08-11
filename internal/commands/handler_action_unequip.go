package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func commandUnequip(cfg *models.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected exactly 1 argument, got %d", len(args))
	}

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	request := models.UnEquipReq{
		Slot: args[0],
	}

	character, err := client.UnEquip(request)
	if err != nil {
		return fmt.Errorf("failed to equip character: %w", err)
	}
	fmt.Printf("Unequipped item from slot %s\n", args[0])

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
