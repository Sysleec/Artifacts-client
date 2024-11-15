package commands

import (
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/artsapi/bot"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func commandBotStart(cfg *models.Config, args ...string) error {
	if len(args) != 3 {
		return fmt.Errorf("expected exactly 3 arguments, got %d", len(args))
	}

	character := args[0]

	if cfg.ApiClient.BotRunning[character] == true {
		return fmt.Errorf("bot for character %s is already running", character)
	}

	client, err := bot.NewClientWrapper(cfg.ApiClient)
	if err != nil {
		return fmt.Errorf("failed to action : %w", err)
	}

	char, err := utils.CheckCharacter(cfg, character)
	if err != nil {
		return fmt.Errorf("failed to check character: %w", err)
	}

	isMaxItems := utils.CheckMaxItems(char)
	if isMaxItems {
		fmt.Printf("character %s has max items\n", character)
		err := utils.BankAllItemsAndReturnToSpot(cfg, char)
		if err != nil {
			return err
		}
	}

	percent := char.MaxHp * 40 / 100

	if char.Hp < percent {
		_, err := client.Client.PostReq("/my/"+char.Name+"/action/rest", []byte{})
		if err != nil {
			return fmt.Errorf("failed to send request: %s", err.Error())
		}

		remainingSeconds := ((char.MaxHp - char.Hp) / 5) + 1
		time.Sleep(time.Duration(remainingSeconds) * time.Second)
	}

	switch args[1] {
	case "gather":
		switch args[2] {
		case "copper":
			err := client.Gather("copper", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "gudgeon":
			err := client.Gather("gudgeon", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "iron":
			err := client.Gather("iron", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "ash_tree":
			err := client.Gather("ash_tree", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "spruce_wood":
			err := client.Gather("spruce_wood", character)
			if err != nil {
				return fmt.Errorf("failed to mining copper: %w", err)
			}
		case "sunflower":
			err := client.Gather("sunflower", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		default:
			return fmt.Errorf("unknown target: %s", args[2])
		}
	case "fighting":
		switch args[2] {
		case "chicken":
			err := client.Fighting("chicken", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "cow":
			err := client.Fighting("cow", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "green_slime":
			err := client.Fighting("green_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "blue_slime":
			err := client.Fighting("blue_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "yellow_slime":
			err := client.Fighting("yellow_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		case "red_slime":
			err := client.Fighting("red_slime", character)
			if err != nil {
				return fmt.Errorf("failed to fighting chicken: %w", err)
			}
		}
	case "craft":
		switch args[2] {
		case "copper_bar":
			err := utils.CraftCopperAndDeposit(cfg, char)
			if err != nil {
				return fmt.Errorf("failed to craft copper: %w", err)
			}
		case "iron":
			err := utils.CraftIronAndDeposit(cfg, char)
			if err != nil {
				return fmt.Errorf("failed to craft iron: %w", err)
			}
		case "cooked_chicken":
			err := utils.CraftAndDeposit(cfg, char, models.CraftAndDeposit{
				Resource:            "raw_chicken",
				Result:              "cooked_chicken",
				QtyResourceForCraft: 1,
				Coords: models.MoveReq{
					X: 1,
					Y: 1,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to craft cooked chicken: %w", err)
			}
		case "cooked_beef":
			err := utils.CraftAndDeposit(cfg, char, models.CraftAndDeposit{
				Resource:            "raw_beef",
				Result:              "cooked_beef",
				QtyResourceForCraft: 1,
				Coords: models.MoveReq{
					X: 1,
					Y: 1,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to craft cooked beef: %w", err)
			}
		}
	default:
		return fmt.Errorf("unknown action: %s", args[1])
	}

	fmt.Printf("Bot for character %s started\n", character)

	return nil
}
