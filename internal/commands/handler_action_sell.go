package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/ge"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"strconv"
	"strings"
	"time"
)

func commandSell(cfg *models.Config, args ...string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected exactly 2 arguments, got %d", len(args))
	}

	target := strings.ToLower(args[0])

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

	qty, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("qty must be an integer: %w", err)
	}

	geItem, err := geClient.GetItem(target)
	if err != nil {
		return fmt.Errorf("failed to get Grand Exchange item : %w", err)
	}

	price := geItem.Data.SellPrice

	request := models.SellReq{
		Code:     target,
		Quantity: qty,
		Price:    price,
	}

	character, err := client.Sell(request)
	if err != nil {
		return fmt.Errorf("failed to sell: %w", err)
	}

	fmt.Printf("Sold %d %s for %d gold\n", qty, target, price)

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
