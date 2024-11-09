package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Sysleec/Artifacts-client/internal/models"
	"github.com/Sysleec/Artifacts-client/internal/utils"
	"time"
)

func (c *ClientWrapper) Gather(resource, char string) error {
	var coords models.MoveReq

	switch resource {
	case "copper":
		coords = models.MoveReq{
			X: 2,
			Y: 0,
		}
	case "gudgeon":
		coords = models.MoveReq{
			X: 4,
			Y: 2,
		}
	case "iron":
		coords = models.MoveReq{
			X: 1,
			Y: 7,
		}
	}

	err := c.goToSpot(coords, char)
	if err != nil {
		return fmt.Errorf("failed to go to %s: %w", resource, err)
	}

	c.Client.BotRunning[char] = true

	go c.gather(char)

	return nil
}

func (c *ClientWrapper) goToResource(coords models.MoveReq, char string) error {
	req := models.MoveReq{
		X: coords.X,
		Y: coords.Y,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.PostReq("/my/"+char+"/action/move", body)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	var action models.Action
	err = json.Unmarshal(resp, &action)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	fmt.Println("Moving to resource...")

	time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)

	return nil
}

func (c *ClientWrapper) gather(char string) {
	for c.Client.BotRunning[char] {
		resp, err := c.Client.PostReq("/my/"+char+"/action/gathering", []byte{})
		if err != nil {
			fmt.Printf("failed to send request: %s", err.Error())
			return
		}

		var action models.Action
		err = json.Unmarshal(resp, &action)
		if err != nil {
			fmt.Printf("failed to unmarshal response: %s", err.Error())
			return
		}

		time.Sleep(time.Duration(action.Data.Cooldown.TotalSeconds) * time.Second)

		isMaxItems := utils.CheckMaxItems(models.ConvertToModelCharacter(action))
		if isMaxItems {
			err := utils.BankAllItemsAndReturnToSpot(&models.Config{ApiClient: c.Client}, models.ConvertToModelCharacter(action))
			if err != nil {
				fmt.Printf("failed to bank all items: %s", err.Error())
				return
			}
		}
	}
	fmt.Printf("\rGather bot for %s stopped              \n", char)
}
