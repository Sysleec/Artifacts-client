package utils

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/action"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/ge"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"time"
)

func CraftIronAndDeposit(cfg *models.Config, char models.Character) error {
	for i := 0; i < 10; i++ {
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

		time.Sleep(time.Duration(secondsRemaining) * time.Second)
		time.Sleep(time.Second)

		fmt.Print("\rCooldown complete!              \n")

		geClient, err := ge.NewClientWrapper(cfg.ApiClient)
		if err != nil {
			return fmt.Errorf("failed to action : %w", err)
		}

		char, err := CheckCharacter(cfg, char.Name)
		if err != nil {
			return fmt.Errorf("failed to check character: %w", err)
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

		//withdraw copper from bank

		bar := char.InventoryMaxItems / 8
		ore := bar * 8

		bankTrans, err := geClient.WithdrawItem(models.BankReq{
			Code:     "iron_ore",
			Quantity: ore,
		})
		if err != nil {
			return err
		}

		secondsRemaining = bankTrans.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		time.Sleep(time.Duration(secondsRemaining) * time.Second)
		time.Sleep(time.Second)

		move, err := client.Move(models.MoveReq{
			X: 1,
			Y: 5,
		})
		if err != nil {
			return err
		}

		secondsRemaining = move.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		time.Sleep(time.Duration(secondsRemaining) * time.Second)
		time.Sleep(time.Second)

		request := models.CraftReq{
			Code:     "iron",
			Quantity: bar,
		}

		character, err := client.Craft(request)
		if err != nil {
			return fmt.Errorf("failed to craft: %w", err)
		}

		fmt.Printf("Crafted %d %s\n", bar, "copper")

		secondsRemaining = character.Data.Cooldown.TotalSeconds
		fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

		time.Sleep(time.Duration(secondsRemaining) * time.Second)
		time.Sleep(time.Second)
	}

	return nil
}
