package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/ge"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func BankAllItemsAndReturnToSpot(cfg *models.Config, char models.Character) error {
	cfg.ApiClient.Character = char.Name
	client, err := action.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	var actionMove models.Action

	act, err := client.Move(models.MoveReq{
		X: 4,
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

	for _, item := range char.Inventory {
		if item.Code == "tasks_coin" || item.Code == "Tasks Coin" || item.Quantity == 0 {
			continue
		}
		bankTrans, err := geClient.DepositItem(models.BankDepositReq{
			Code:     item.Code,
			Quantity: item.Quantity,
		})
		if err != nil {
			return fmt.Errorf("failed to get Grand Exchange item : %w", err)
		}

		secondsRemaining = bankTrans.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		for secondsRemaining > 0 {
			fmt.Printf("\rTime left: %d seconds", secondsRemaining)
			time.Sleep(1 * time.Second)
			secondsRemaining--
		}

		fmt.Print("\rCooldown complete!              \n")
	}

	fmt.Printf("Banked all items\n")

	move, err := client.Move(models.MoveReq{
		X: char.X,
		Y: char.Y,
	})
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(move.Data.Cooldown.RemainingSeconds) * time.Second)

	return nil
}
