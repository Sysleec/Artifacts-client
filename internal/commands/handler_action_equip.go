package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func commandEquip(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	request := models.EquipReq{
		Code: args[0],
		Slot: args[1],
	}

	requestUnEquip := models.UnEquipReq{
		Slot: args[1],
	}

	character, _ := client.UnEquip(requestUnEquip)

	fmt.Printf("Unequipped item from slot %s\n", args[1])

	secondsRemaining := character.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	character, err = client.Equip(request)
	if err != nil {
		return fmt.Errorf("failed to equip character: %w", err)
	}
	fmt.Printf("Equipped item %s to slot %s\n", args[0], args[1])

	secondsRemaining = character.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	return nil
}
