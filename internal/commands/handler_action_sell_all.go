package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/ge"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func commandSellAll(cfg *models.Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("expected exactly 0 arguments, got %d", len(args))
	}

	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	var actionMove models.Action

	act, err := client.Move(models.MoveReq{
		X: 5,
		Y: 1,
	})
	if err != nil {
		return err
	}
	actionMove = act

	fmt.Printf("Moved character to x = %d, y = %d\n", actionMove.Data.Destination.X, actionMove.Data.Destination.Y)

	secondsRemaining := actionMove.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	geClient, err := ge.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	char, err := utils.CheckCharacter(cfg, cfg.ApiClient.Character)
	if err != nil {
		return fmt.Errorf("failed to check character: %w", err)
	}

	for _, item := range char.Inventory {
		geItem, err := geClient.GetItem(item.Code)
		if err != nil {
			return fmt.Errorf("failed to get Grand Exchange item : %w", err)
		}

		price := geItem.Data.SellPrice

		request := models.SellReq{
			Code:     item.Code,
			Quantity: item.Quantity,
			Price:    price,
		}

		character, err := client.Sell(request)
		if err != nil {
			return fmt.Errorf("failed to sell: %w", err)
		}

		fmt.Printf("Sold %d %s for %d gold\n", item.Quantity, item.Code, price)

		secondsRemaining = character.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		for secondsRemaining > 0 {
			fmt.Printf("\rTime left: %d seconds", secondsRemaining)
			time.Sleep(1 * time.Second)
			secondsRemaining--
		}

		fmt.Print("\rCooldown complete!              \n")
	}

	fmt.Printf("Sold all items\n")

	return nil
}
