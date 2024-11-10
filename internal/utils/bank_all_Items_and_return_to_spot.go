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

	secondsRemaining := char.Cooldown
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	time.Sleep(time.Duration(secondsRemaining) * time.Second)
	time.Sleep(time.Second)

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

	secondsRemaining = actionMove.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	time.Sleep(time.Duration(secondsRemaining) * time.Second)
	time.Sleep(time.Second)

	fmt.Print("\rCooldown complete!              \n")

	geClient, err := ge.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	for _, item := range char.Inventory {
		if item.Code == "tasks_coin" || item.Code == "Tasks Coin" || item.Quantity == 0 {
			continue
		}

		bankTrans, err := geClient.DepositItem(models.BankReq{
			Code:     item.Code,
			Quantity: item.Quantity,
		})
		if err != nil {
			return fmt.Errorf("failed to deposit Grand Exchange item : %w", err)
		}

		secondsRemaining = bankTrans.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		time.Sleep(time.Duration(secondsRemaining) * time.Second)
		time.Sleep(time.Second)

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

	secondsRemaining = move.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	time.Sleep(time.Duration(secondsRemaining) * time.Second)
	time.Sleep(time.Second)

	return nil
}
