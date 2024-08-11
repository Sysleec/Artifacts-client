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

	var actionMove models.Action

	switch request.Code {
	case "copper":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 5,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "copper_ring":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 3,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "copper_dagger":
		act, err := client.Move(models.MoveReq{
			X: 2,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "sticky_dagger":
		act, err := client.Move(models.MoveReq{
			X: 2,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "copper_boots":
		act, err := client.Move(models.MoveReq{
			X: 3,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "copper_helmet":
		act, err := client.Move(models.MoveReq{
			X: 3,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "copper_legs_armor":
		act, err := client.Move(models.MoveReq{
			X: 3,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "feather_coat":
		act, err := client.Move(models.MoveReq{
			X: 3,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "body_armor":
		act, err := client.Move(models.MoveReq{
			X: 3,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "cooked_chicken":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "cooked_gudgeon":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "cooked_beef":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "cooked_shrimp":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "cheese":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "fried_eggs":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "mushroom_soup":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	case "beef_stew":
		act, err := client.Move(models.MoveReq{
			X: 1,
			Y: 1,
		})
		if err != nil {
			return err
		}
		actionMove = act
	default:
		return fmt.Errorf("unknown target: %s", request.Code)
	}

	fmt.Printf("Moved character to x = %d, y = %d\n", actionMove.Data.Destination.X, actionMove.Data.Destination.Y)

	secondsRemaining := actionMove.Data.Cooldown.TotalSeconds
	fmt.Printf("Waiting for %d seconds\n", secondsRemaining)

	for secondsRemaining > 0 {
		fmt.Printf("\rTime left: %d seconds", secondsRemaining)
		time.Sleep(1 * time.Second)
		secondsRemaining--
	}

	fmt.Print("\rCooldown complete!              \n")

	character, err := client.Craft(request)
	if err != nil {
		return fmt.Errorf("failed to craft: %w", err)
	}

	fmt.Printf("Crafted %d %s\n", qty, target)

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
